<template>
  <div>
    <t-drawer
      header="附件详情"
      :visible="visible"
      :onClose="close"
      :closeBtn="true"
      size="480px"
      :zIndex="2600"
      :onOverlayClick="close"
      placement="right"
    >
      <t-list :split="true" v-if="attachment">
        <t-list-item>
          <t-list-item-meta title="附件名"></t-list-item-meta>
          {{ attachment.fileName }}
        </t-list-item>
        <t-list-item>
          <t-list-item-meta title="存储位置"></t-list-item-meta>
          {{ attachmentStorageType }}
        </t-list-item>
        <t-list-item>
          <t-list-item-meta title="图片大小"></t-list-item-meta>
          {{ attachment.size }}
        </t-list-item>
        <t-list-item>
          <t-list-item-meta title="图片尺寸"></t-list-item-meta>
          {{ attachment.height }} x {{ attachment.width }}
        </t-list-item>
        <t-list-item>
          <t-list-item-meta title="上传日期"></t-list-item-meta>
          {{ attachment.uploadTime }}
        </t-list-item>
        <t-list-item>
          <t-list-item-meta title="图片链接"></t-list-item-meta>
          {{ attachment.url }}
          <template #action>
            <t-button variant="text" shape="square" @click="copyURL"
              ><file-copy-icon
            /></t-button>
          </template>
        </t-list-item>
        <t-list-item>
          <t-list-item-meta title="Markdown 链接"></t-list-item-meta>
          ![{{ attachment.fileName }}]({{ attachment.url }})
          <template #action>
            <t-button variant="text" shape="square" @click="copyMarkdown"
              ><file-copy-icon
            /></t-button>
          </template>
        </t-list-item>
      </t-list>
      <template #footer>
        <t-button @click="deleteAttach">删除</t-button>
      </template>
    </t-drawer>
  </div>
</template>

<script>
import { FileCopyIcon } from "tdesign-icons-vue";
export default {
  name: "AttachDetailDrawer",
  data() {
    return {
      attachment: {},
      visible: false,
    };
  },
  computed: {
    attachmentStorageType() {
      switch (this.attachment.type) {
        case "local":
          return "本地";
        case "qcloud_cos":
          return "腾讯云 COS";
        case "aliyun_oss":
          return "阿里云 OSS";
        case "huawei_obs":
          return "华为 OBS";
        case "qiniuyun":
          return "七牛云";
        case "youpaiyun":
          return "又拍云";
        default:
          return "未知";
      }
    },
  },
  methods: {
    open(attachment) {
      this.visible = true;
      this.attachment = attachment;
    },
    close() {
      this.visible = false;
    },
    deleteAttach() {
      this.$message.info("未实现");
    },
    copyURL() {
      this.$copyText(this.attachment.url);
      this.$message.success("已复制");
    },
    copyMarkdown() {
      this.$copyText(`![${this.attachment.fileName}](${this.attachment.url})`);
      this.$message.success("已复制");
    },
  },
  components: { FileCopyIcon },
};
</script>

<style lang="less" scoped>
.image-wrapper {
  margin: 0 auto;
  justify-content: center;
  img {
    height: 175px;
    align-items: center;
    margin: 0 auto;
    border-radius: 2px;
  }
}
</style>