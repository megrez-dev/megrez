<template>
  <!-- 不能用骨架屏包裹vditor，因为包裹之后vditor挂载时会找不到#vditor元素 -->
  <div id="vditor">
    <!-- 因此骨架屏单独写在此处，仅用于占位，当vditor挂在后会替换它，无需关心loading何时结束 -->
    <t-skeleton
      class="vditor-skeleton"
      animation="gradient"
      :loading="loading"
      :rowCol="skeletonRowCol"
      :delay="delayTime"
    ></t-skeleton>
  </div>
</template>

<script>
import Vditor from 'vditor';
import { VDITOR_BASE_CONFIGS, SKELETON_ROW_COL, THEME_SETTING, VDITOR_TOOLBAR, INSERT_IMG_CONFIGS } from '@/views/article/constants';
import 'vditor/dist/index.css';

export default {
  props: {
    value: {
      type: String,
      require: false,
      default: '',
    },
    loading: Boolean,
  },
  computed: {
    isDark() {
      return this.$store.state.app.isDark;
    },
  },
  watch: {
    // 监听 store里面的数据
    isDark(isDark) {
      const { light, dark } = THEME_SETTING;
      isDark
        ? this.contentEditor?.setTheme(dark.theme, dark.contentTheme)
        : this.contentEditor?.setTheme(light.theme, light.contentTheme);
    },
  },
  data() {
    return {
      contentEditor: null,
      delayTime: 500,
      skeletonRowCol: SKELETON_ROW_COL,
    };
  },
  methods: {
    contentChange() {
      this.$emit('input', this.contentEditor.getValue());
    },
    countWord(length) {
      this.$emit('countWord', length);
    },
    append(content) {
      this.contentEditor.insertValue(content);
    },
    initViditor(cb) {
      const { isDark, value } = this;
      const { light, dark } = THEME_SETTING;
      const initTheme = isDark ? dark : light;
      this.contentEditor = new Vditor('vditor', {
        ...VDITOR_BASE_CONFIGS,
        input: this.contentChange,
        value,
        theme: initTheme.theme,
        toolbar: this.initToolbar(),
        preview: {
          delay: 50,
          theme: {
            current: initTheme.contentTheme,
          }
        },
        counter: {
          enable: true,
          type: 'text',
          after: (length) => {
            this.countWord(length);
            cb();
          },
        },
      });
    },
    initToolbar() {
      // 创建新的数组，防止直接修改VDITOR_TOOLBAR
      const toolbar = [...VDITOR_TOOLBAR];
      const indexOfLink = toolbar.indexOf('link');
      toolbar.splice(indexOfLink, 0, {
        ...INSERT_IMG_CONFIGS,
        click: () => {
          this.$emit('insertImage');
        }
      });
      return toolbar;
    }
  },
};
</script>

<style lang="less" scoped>
#vditor {
  height: calc(100vh - 250px);
  z-index: 9999 !important;
}
.vditor-skeleton {
  padding: 20px;
  height: 100%;
}
</style>
