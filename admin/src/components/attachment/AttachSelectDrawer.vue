<template>
  <div>
    <t-drawer header="附件列表" :visible="visible" :onClose="close" :closeBtn="true" size="400px" :zIndex="2600"
      :onOverlayClick="close" placement="right">
      <t-row :gutter="[16, 16]">
        <t-col>
          <t-input v-model="keywords" placeholder="通过关键词搜索"></t-input>
        </t-col>
      </t-row>
      <t-skeleton class="images-skeleton" animation="gradient" :loading="loading" :delay="300">
        <div class="images-wrapper">
          <m-image fit="cover" v-for="(attach, index) in attachments" class="image-item"
          :class="getIndexInSelected(attach) >= 0 ? 'image-selected' : ''" :src="attach.url" :key="attach.id"
          @maskClick="handleMaskClick(attach, index)">
          <template slot="maskInner">
            <span style="cursor: pointer" @click.stop="previewImage(attach.url)">
              <BrowseIcon size="1.1em" /> 预览
            </span>
            <t-divider layout="vertical"></t-divider>
            <span style="cursor: pointer" @click.stop="previewAttach(attach)">
              <InfoCircleIcon size="1.1em" /> 详情
            </span>
          </template>
        </m-image>
      </div>
      </t-skeleton>
      <t-pagination :current="pagination.pageNum" :total="pagination.total" :pageSizeOptions="[]"
        @change="onPaginationChange"></t-pagination>
      <template #footer>
        <t-button @click="onClickUpload">
          <UploadIcon slot="icon" />
          上传附件
        </t-button>
        <span class="complete-button" v-if="mode === 'multiple'">
          <t-button theme="danger" :disabled="selectedNum <= 0" @click="onClickComplete">
            完成{{ maxNum > 0 ? `(${selectedNum}/${maxNum})` : '' }}
          </t-button>
        </span>
      </template>
    </t-drawer>
    <AttachDetailDrawer ref="attachDetailDrawer"></AttachDetailDrawer>
    <AttachUploadDialog ref="attachUploadDialog" @uploadSuccess="syncData"></AttachUploadDialog>
  </div>
</template>

<script>
import AttachUploadDialog from '@/components/attachment/AttachUploadDialog.vue';
import AttachDetailDrawer from '@/components/attachment/AttachDetailDrawer.vue';
import MImage from '@/components/image/Image.vue';
import { viewerOptions } from '@/components/image/constants';
import { UploadIcon, BrowseIcon, InfoCircleIcon } from 'tdesign-icons-vue';

export default {
  name: 'AttachSelectDrawer',
  components: {
    AttachUploadDialog,
    AttachDetailDrawer,
    MImage,
    UploadIcon,
    BrowseIcon,
    InfoCircleIcon,
  },
  data() {
    return {
      hasInit: false,
      keywords: '',
      attachments: [],
      loading: false,
      selectedAttaches: [],
      pagination: {
        pageSize: 10,
        pageNum: 1,
        total: 0,
      },
    };
  },
  props: {
    mode: {
      type: String,
      required: false,
      default: 'single',
    },
    // AttachSelectDrawer现在为受控组件，是否展示受控于父组件传递的visible属性
    visible: {
      type: Boolean,
      required: false,
      default: false,
    },
    maxNum: {
      type: [Number, String],
      default: 0,
    },
  },
  computed: {
    selectedNum() {
      return this.selectedAttaches.length;
    },
  },
  watch: {
    visible(newVal) {
      if (newVal) {
        this.$emit('open');
        // 若未初始化，则请求数据，进行数据的初始化
        if (this.hasInit) return;
        this.hasInit = true;
        this.syncData();
      } else {
        this.$emit('close');
      }
    },
  },
  methods: {
    close() {
      // 同步更新visible值，需要父组件传参时添加.sync修饰符
      this.$emit('update:visible', false);
      this.$emit('close');
      this.selectedAttaches = [];
    },
    onClickUpload() {
      this.$refs.attachUploadDialog.open();
    },
    onPaginationChange(pagination) {
      if (pagination) {
        const { current: pageNum, pageSize } = pagination;
        this.pagination = { ...this.pagination, pageNum, pageSize };
      }
      this.syncData();
    },
    syncData() {
      this.loading = true;
      const { pageNum, pageSize } = this.pagination;
      this.$request
        .get('attachments', { params: { pageNum, pageSize } })
        .then((res) => {
          if (res.status === 0) {
            this.pagination.total = res.data.total;
            this.attachments = res.data.list || [];
          } else {
            this.$message.error('获取附件列表失败');
          }
        })
        .catch(() => {
          this.$message.error('获取附件列表失败');
        })
        .finally(() => {
          this.loading = false;
        });
    },
    previewAttach(attach) {
      this.$refs.attachDetailDrawer.open(attach);
    },
    previewImage(url) {
      this.$viewerApi({
        images: [url],
        options: viewerOptions,
      });
    },
    handleMaskClick(attach) {
      const indexInSelected = this.getIndexInSelected(attach);
      // 如果attach在selectedAttaches中，说明已经被选中，则执行取消选择操作
      indexInSelected >= 0 ? this.unSelect(indexInSelected) : this.select(attach);
    },
    getIndexInSelected(attach) {
      return this.selectedAttaches.findIndex((selectedAttach) => selectedAttach.id === attach.id);
    },
    select(attach) {
      const maxNum = Number(this.maxNum);
      if (maxNum > 0 && maxNum <= this.selectedNum) {
        this.$message.warning('最多只能选择' + maxNum + '个附件');
        return;
      }
      this.selectedAttaches.push(attach);
      if (this.mode === 'single') {
        this.$emit('select', attach);
        this.close();
        return;
      }
    },
    unSelect(deleteIndex) {
      this.selectedAttaches.splice(deleteIndex, 1);
    },
    onClickComplete() {
      this.$emit(
        'select',
        this.selectedAttaches,
      );
      this.close();
    },
  },
};
</script>

<style lang="less" scoped>
@image-wrapper-height: calc(100% - 84px);
.images-skeleton {
  width: 100%;
  margin: 10px 0px;
  height: @image-wrapper-height;
}
.images-wrapper {
  position: relative;
  display: flex;
  align-content: flex-start;
  flex-wrap: wrap;
  height: 100%;
  width: 100%;
  margin: 10px 0px;
  overflow: auto;
  .image-item {
    position: relative;
    flex: 1 1 160px;
    height: 170px;
    margin: 2px;
    overflow: hidden;
    background-color: @bg-color-secondarycontainer;
    border: 5px solid var(--td-bg-color-container);
    &.image-selected {
      border: 5px solid @brand-color-6;
    }
  }
}

.complete-button {
  float: right;
}
</style>
