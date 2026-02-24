<template>
  <PageView>
    <template slot="header">
      <div class="page-header-bar">
        <div class="page-header-bar-text">
          <span class="page-header-bar-title">基础设置</span>
        </div>
        <span class="page-header-bar-operator">
          <span class="page-header-bar-operator-item">
            <t-button
              theme="primary"
              variant="base"
              @click="onClickSave"
              :loading="saveBtnLoading"
              >保存设置</t-button
            >
          </span>
        </span>
      </div>
    </template>
    <template slot="content">
      <div class="settings-container">
        <t-tabs defaultValue="basic">
          <t-tab-panel value="basic">
            <template #label>
              <icon name="setting" style="margin-right: 4px" /> 基本设置
            </template>
            <div class="settings-form-container">
              <t-form ref="form" labelAlign="top" :colon="true">
                <t-form-item label="博客标题" name="blogTitle">
                  <t-input v-model="settings.basic.blogTitle"></t-input>
                </t-form-item>
                <t-form-item label="博客简介" name="blogDescription">
                  <t-textarea
                    v-model="settings.basic.blogDescription"
                    placeholder="一句简短的话描述你的博客"
                    :autosize="{ minRows: 2 }"
                  />
                </t-form-item>
                <t-form-item label="博客地址" name="blogURL">
                  <t-input
                    v-model="settings.basic.blogURL"
                    placeholder="http://"
                  ></t-input>
                </t-form-item>
                <t-form-item label="Favicon" name="blogFavicon">
                  <t-input-group separate>
                    <t-input
                      v-model="settings.basic.blogFavicon"
                      clearable
                      :style="{ width: '500px' }"
                    ></t-input>
                    <t-button
                      theme="primary"
                      shape="square"
                      variant="outline"
                      @click="openFaviconSelectDrawer()"
                      ><image-icon slot="icon"
                    /></t-button>
                  </t-input-group>
                </t-form-item>
              </t-form>
            </div>
          </t-tab-panel>
          <t-tab-panel value="attachment">
            <template #label>
              <icon name="image" style="margin-right: 4px" /> 附件设置
            </template>
            <div class="settings-form-container">
              <t-form ref="form" labelAlign="top" :colon="true">
                <t-form-item label="上传位置">
                  <t-select
                    v-model="settings.attachment.uploadType"
                    :options="uploadTypeOptions"
                  />
                </t-form-item>
                <t-form-item
                  label="secretID"
                  v-if="settings.attachment.uploadType === 'qcloud_cos'"
                >
                  <t-input
                    v-model="settings.attachment.qCloudCos.secretID"
                  ></t-input>
                </t-form-item>
                <t-form-item
                  label="secretKey"
                  v-if="settings.attachment.uploadType === 'qcloud_cos'"
                >
                  <t-input
                    v-model="settings.attachment.qCloudCos.secretKey"
                  ></t-input>
                </t-form-item>
                <t-form-item
                  label="域名"
                  v-if="settings.attachment.uploadType === 'qcloud_cos'"
                >
                  <t-input
                    v-model="settings.attachment.qCloudCos.domain"
                  ></t-input>
                </t-form-item>
                <t-form-item
                  label="上传路径"
                  v-if="settings.attachment.uploadType === 'qcloud_cos'"
                >
                  <t-input
                    v-model="settings.attachment.qCloudCos.path"
                    placeholder="如：blog/images"
                  ></t-input>
                </t-form-item>
                <t-form-item>
                  <t-button
                    @click="pingQcloudCos"
                    v-if="settings.attachment.uploadType === 'qcloud_cos'"
                    >测试连接</t-button
                  >
                </t-form-item>
              </t-form>
            </div>
          </t-tab-panel>
        </t-tabs>
        <AttachSelectDrawer
          ref="faviconSelectDrawer"
          mode="single"
          :visible.sync="faviconSelectDrawerVisiable"
          @select="selectFavicon"
        ></AttachSelectDrawer>
      </div>
    </template>
  </PageView>
</template>
<script>
import { Icon, ImageIcon } from "tdesign-icons-vue";
import AttachSelectDrawer from "@/components/attachment/AttachSelectDrawer.vue";
import PageView from "@/layouts/PageView";
export default {
  name: "Settings",
  data() {
    return {
      saveBtnLoading: false,
      faviconSelectDrawerVisiable: false,
      settings: {
        basic: {
          blogTitle: "",
          blogDescription: "",
          blogURL: "",
          blogFavicon: "",
        },
        attachment: {
          uploadType: "local",
          qCloudCos: {
            secretID: "",
            secretKey: "",
            domain: "",
            path: "",
          },
        },
      },
      uploadTypeOptions: [
        {
          value: "local",
          label: "本地",
        },
        {
          value: "qcloud_cos",
          label: "腾讯云 COS",
        },
        // TODO: 添加其他上传方式
        // {
        //   value: "aliyun_oss",
        //   label: "阿里云 OSS",
        // },
        // {
        //   value: "huawei_obs",
        //   label: "华为云 OBS",
        // },
        // {
        //   value: "qiniuyun",
        //   label: "七牛云",
        // },
        // {
        //   value: "youpaiyun",
        //   label: "又拍云",
        // },
      ],
    };
  },
  methods: {
    pingQcloudCos() {
      // TODO: ping 期间将按钮设置为 loading 状态
      this.$request
        .post("upload/qcloudcos/ping", this.settings.attachment.qCloudCos)
        .then(() => {
          this.$message.success("连接成功");
        });
    },
    onClickSave() {
      this.saveBtnLoading = true;
      this.$request
        .put("settings", this.settings)
        .then(() => {
          this.$message.success("保存成功");
        })
        .finally(() => {
          this.saveBtnLoading = false;
        });
    },
    fetchSettings() {
      this.$request.get("settings").then((res) => {
        this.settings = res.data;
      });
    },
    openFaviconSelectDrawer() {
      this.faviconSelectDrawerVisiable = true;
    },
    selectFavicon(attach) {
      this.settings.basic.blogFavicon = attach.url;
    },
  },
  beforeMount() {
    this.fetchSettings();
  },
  components: {
    Icon,
    ImageIcon,
    AttachSelectDrawer,
    PageView,
  },
};
</script>

<style lang="less" scoped>
.page-header-bar {
  margin-bottom: 10px;
  display: flex;
  justify-content: space-between;
  .page-header-bar-text {
    display: flex;
    flex-direction: row;
    align-items: center;
    .page-header-bar-title {
      font-size: 20px;
      font-weight: bold;
      text-align: center;
      align-items: center;
      color: @text-color-primary;
    }
    .page-header-bar-description {
      margin-left: 10px;
      font-size: 16px;
      font-weight: normal;
      text-align: center;
      align-items: center;
      color: @text-color-secondary;
    }
  }
  .page-header-bar-operator {
    float: right;
    .page-header-bar-operator-item {
      margin-left: 15px;
    }
  }
}
.settings-container {
  .settings-form-container {
    padding: 30px 30px;
    max-width: 600px;
  }
}
</style>