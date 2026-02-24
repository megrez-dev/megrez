<template>
  <PageView>
    <template slot="header">
      <div class="page-header-bar">
        <div class="page-header-bar-text">
          <span class="page-header-bar-title">当前主题</span>
          <span class="page-header-bar-description">{{ currentTheme }}</span>
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
      <div class="theme-setting-container">
        <t-tabs
          :defaultValue="themeConfig.tabs[0].key"
          v-if="themeConfig.tabs.length != 0"
        >
          <t-tab-panel
            :value="tab.key"
            v-for="tab in themeConfig.tabs"
            :key="tab.key"
          >
            <template #label>
              <icon name="setting" style="margin-right: 4px" /> {{ tab.name }}
            </template>
            <div class="setting-form-container">
              <t-form ref="form" labelAlign="top" :colon="true">
                <t-form-item
                  :label="item.name"
                  :name="item.key"
                  :help="item.description"
                  v-for="item in tab.items"
                  :key="item.key"
                >
                  <t-input
                    v-model="item.value"
                    :placeholder="item.placeholder"
                    v-if="item.type === 'input'"
                  ></t-input>
                  <t-textarea
                    v-model="item.value"
                    :placeholder="item.placeholder"
                    :name="item.key"
                    :autosize="{ minRows: 6 }"
                    v-if="item.type === 'textarea'"
                  />
                  <t-select
                    v-model="item.value"
                    :options="item.options"
                    v-if="item.type === 'select'"
                  />
                  <t-select
                    v-model="item.value"
                    multiple
                    :minCollapsedNum="3"
                    :options="item.options"
                    v-if="item.type === 'multiSelect'"
                  />
                  <t-switch
                    v-model="item.value"
                    v-if="item.type === 'switch'"
                  ></t-switch>
                  <t-tag-input
                    v-model="item.value"
                    clearable
                    v-if="item.type === 'tags'"
                  />
                  <!-- <t-input
                v-model="item.values"
                :placeholder="item.placeholder"
                v-if="item.type === 'tags'"
              ></t-input> -->

                  <t-input-group separate v-if="item.type === 'image'">
                    <t-input
                      v-model="item.value"
                      clearable
                      :placeholder="item.placeholder"
                      :style="{ width: '500px' }"
                    ></t-input>
                    <t-button
                      theme="primary"
                      shape="square"
                      variant="outline"
                      @click="openAttachSelectDrawer(tab, item)"
                      ><image-icon slot="icon"
                    /></t-button>
                  </t-input-group>
                </t-form-item>
              </t-form>
            </div>
          </t-tab-panel>
        </t-tabs>
        <AttachSelectDrawer
          ref="attachSelectDrawer"
          mode="single"
          :visible.sync="drawerVisible"
          @select="selectAttach"
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
  name: "ThemeSetting",
  data() {
    return {
      saveBtnLoading: false,
      themeConfig: {
        tabs: [],
      },
      drawerVisible: false,
      currentTheme: "",
      selectedAttachTabIndex: 0,
      selectedAttachItemIndex: 0,
    };
  },
  methods: {
    openAttachSelectDrawer(tab, item) {
      this.drawerVisible = true;
      this.selectedAttachTabIndex = this.themeConfig.tabs.indexOf(tab);
      this.selectedAttachItemIndex =
        this.themeConfig.tabs[this.selectedAttachTabIndex].items.indexOf(item);
    },
    onClickSave() {
      this.saveBtnLoading = true;
      this.$request
        .put("/theme/current/config", this.themeConfig)
        .then(() => {
          this.$message.success("保存成功");
        })
        .finally(() => {
          this.saveBtnLoading = false;
        });
    },
    fetchConfig() {
      this.$request.get("/theme/current/config").then((res) => {
        for (let i = 0; i < this.themeConfig.tabs.length; i++) {
          for (let j = 0; j < this.themeConfig.tabs[i].items.length; j++) {
            if (
              this.themeConfig.tabs[i].items[j].type === "multiSelect" ||
              this.themeConfig.tabs[i].items[j].type === "tags"
            ) {
              this.themeConfig.tabs[i].items[j].value = [];
            }
          }
        }
        this.themeConfig = res.data;
      });
    },
    fetchTheme() {
      this.$request.get("/theme/current/id").then((res) => {
        this.currentTheme = res.data;
      });
    },
    selectAttach(attach) {
      this.themeConfig.tabs[this.selectedAttachTabIndex].items[
        this.selectedAttachItemIndex
      ].value = attach.url;
    },
  },
  beforeMount() {
    this.fetchConfig();
    this.fetchTheme();
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
.theme-setting-container {
  .setting-form-container {
    padding: 30px 30px;
    max-width: 600px;
  }
}
</style>