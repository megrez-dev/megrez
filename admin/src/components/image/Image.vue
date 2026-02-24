<template>
  <div
    class="megrez-image"
    :style="$attrs.style"
    @mouseenter="maskVisable = true"
    @mouseleave="maskVisable = false"
  >
    <div class="megrez-image-loading" v-if="loading"><LoadingIcon /></div>
    <img
      class="megrez-image-img"
      :src="error ? fallbackSrc : src"
      :alt="alt"
      :style="imageStyle"
    />
    <div
      :style="{ opacity: maskVisable ? 1 : 0, cursor: $slots.maskInner ? 'default' : 'pointer' }"
      class="megrez-image-mask"
      @click="handleMaskClick"
    >
      <!-- 父组件可以通过slot定义mask内展示内容 -->
      <span v-if="!$slots.maskInner && preview" ><BrowseIcon />{{ typeof preview === 'string' ?  preview : '预览' }}</span>
      <slot v-else name="maskInner"></slot>
    </div>
  </div>
</template>

<script>
import { BrowseIcon, LoadingIcon } from 'tdesign-icons-vue';
import { viewerOptions, defaultImageUrl } from '@/components/image/constants';

const isSupportObjectFit = () => document.documentElement.style.objectFit !== undefined;
const ObjectFit = {
  NONE: 'none',
  CONTAIN: 'contain',
  COVER: 'cover',
  FILL: 'fill',
  SCALE_DOWN: 'scale-down',
};
export default {
  name: 'MegrezImage',
  components: {
    LoadingIcon,
    BrowseIcon,
  },
  props: {
    src: String,
    alt: String,
    fit: String,
    fallback: String,
    preview: {
      type: [ Boolean, String ],
      default: true,
    },
    // 用于配置v-viewer
    viewerOptions: {
      type: Object,
      required: false,
      default: () => {
        return {};
      },
    },
    maskClass: {
      type: String || Array,
      required: false,
      default: '',
    },
  },
  data() {
    return {
      error: false,
      loading: true,
      imageWidth: 0,
      imageHeight: 0,
      maskVisable: false,
    };
  },
  computed: {
    imageStyle() {
      const { fit } = this;
      if (fit) {
        return isSupportObjectFit() ? { 'object-fit': fit } : this.getImageStyle(fit);
      }
      return {};
    },
    fallbackSrc() {
      return this.fallback || defaultImageUrl;
    },
  },
  mounted() {
    this.loadImage();
  },
  methods: {
    loadImage() {
      this.loading = true;
      this.error = false;
      const img = new Image();
      img.onload = (e) => this.handleLoad(e, img);
      img.onerror = this.handleError.bind(this);
      img.src = this.src;
    },
    handleLoad(e, img) {
      this.imageWidth = img.width;
      this.imageHeight = img.height;
      this.loading = false;
      this.error = false;
      this.$emit('load', e);
    },
    handleError(e) {
      this.loading = false;
      this.error = true;
      this.$emit('error', e);
    },
    handleMaskClick(e) {
      this.$emit('maskClick', e);
      if (!this.$slots.maskInner && this.preview) this.handlePreview();
    },
    handlePreview() {
      this.$viewerApi({
        images: [this.src],
        options: { ...viewerOptions, ...this.viewerOptions },
      });
    },
    // 参考element-ui，低版本浏览器没有object-fit方法，需要用此兼容
    getImageStyle(fit) {
      const { imageWidth, imageHeight } = this;
      const { clientWidth: containerWidth, clientHeight: containerHeight } = this.$el;

      if (!imageWidth || !imageHeight || !containerWidth || !containerHeight) return {};

      const imageAspectRatio = imageWidth / imageHeight;
      const containerAspectRatio = containerWidth / containerHeight;

      if (fit === ObjectFit.SCALE_DOWN) {
        const isSmaller = imageWidth < containerWidth && imageHeight < containerHeight;
        fit = isSmaller ? ObjectFit.NONE : ObjectFit.CONTAIN;
      }

      switch (fit) {
        case ObjectFit.NONE:
          return { width: 'auto', height: 'auto' };
        case ObjectFit.CONTAIN:
          return imageAspectRatio < containerAspectRatio ? { width: 'auto' } : { height: 'auto' };
        case ObjectFit.COVER:
          return imageAspectRatio < containerAspectRatio ? { height: 'auto' } : { width: 'auto' };
        default:
          return {};
      }
    },
  },
};
</script>

<style lang="less" scoped>
.megrez-image {
  width: 100%;
  height: 100%;
  display: inline-block;
  position: relative;
  .megrez-image-loading {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .megrez-image-img {
    position: absolute;
    width: 100%;
    height: 100%;
    vertical-align: middle;
    z-index: 100;
  }
  .megrez-image-mask {
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    width: 100%;
    height: 100%;
    color: @text-color-anti;
    background-color: @mask-active;
    opacity: 0;
    z-index: 200;
    transition: all 0.2s;
    cursor: pointer;
  }
}
</style>
