<template>
  <PageView>
    <template slot="header">
      <div class="page-header-bar">
        <div class="page-header-bar-text">
          <span class="page-header-bar-title">当前主题</span>
          <span class="page-header-bar-description">{{
            currentTheme.name
          }}</span>
        </div>
        <span class="page-header-bar-operator">
          <span class="page-header-bar-operator-item">
            <t-button theme="primary" variant="base" @click="onClickInstall"
              >安装</t-button
            >
          </span>
        </span>
      </div>
    </template>
    <template slot="content">
      <div class="themes-container">
        <div class="card-container" v-for="theme in themes" :key="theme.id">
          <t-card :title="theme.name" :cover="theme.cover">
            <template #footer>
              <t-row :align="'middle'" justify="center">
                <t-col flex="auto">
                  <t-button
                    block
                    variant="text"
                    v-if="theme.isCurrent"
                    disabled
                  >
                    <lock-on-icon slot="icon" style="color: blue" />已启用
                  </t-button>
                  <t-button
                    block
                    variant="text"
                    v-else
                    @click="onClickApply(theme.id)"
                  >
                    <lock-off-icon slot="icon" />未启用
                  </t-button>
                </t-col>
                <t-col flex="none">
                  <t-divider layout="vertical" />
                </t-col>
                <t-col flex="auto">
                  <t-button
                    block
                    variant="text"
                    v-if="theme.isCurrent"
                    @click="onClickSetting"
                  >
                    <setting-icon slot="icon" />设置
                  </t-button>
                  <t-button block variant="text" disabled v-else>
                    <setting-icon slot="icon" />设置
                  </t-button>
                </t-col>
                <t-col flex="none">
                  <t-divider layout="vertical" />
                </t-col>
                <t-col flex="auto">
                  <t-button
                    block
                    variant="text"
                    v-if="theme.isCurrent"
                    disabled
                    @click="onClickDelete(theme.id)"
                  >
                    <delete-icon slot="icon" />删除
                  </t-button>
                  <t-button
                    block
                    variant="text"
                    v-else
                    @click="onClickDelete(theme.id)"
                  >
                    <delete-icon slot="icon" />删除
                  </t-button>
                </t-col>
              </t-row>
            </template>
          </t-card>
        </div>
      </div>
      <ThemeInstallDialog
        ref="themeInstallDialog"
        @installSuccess="onInstallSuccess"
      ></ThemeInstallDialog>
    </template>
  </PageView>
</template>

<script>
import PageView from "@/layouts/PageView";
import ThemeInstallDialog from "@/views/appearance/components/ThemeInstallDialog";
import {
  SettingIcon,
  LockOnIcon,
  LockOffIcon,
  DeleteIcon,
} from "tdesign-icons-vue";
export default {
  name: "Theme",
  data() {
    return {
      themes: [],
      currentThemeID: "",
      currentTheme: {},
    };
  },
  methods: {
    onClickSetting() {
      this.$router.push({
        name: "ThemeSetting",
      });
    },
    onClickApply(value) {
      this.$request
        .put("option/blog_theme", {
          value: value,
        })
        .then(() => {
          this.$message.success("主题已启用");
          this.fetchData();
        });
    },
    onClickDelete(value) {
      this.$request.delete("theme/" + value).then(() => {
        this.$message.success("已删除主题");
        this.fetchData();
      });
    },
    onClickInstall() {
      this.$refs.themeInstallDialog.open();
    },
    onInstallSuccess() {
      this.fetchData();
    },
    fetchData() {
      this.$request
        .get("/theme/current/id")
        .then((res) => {
          this.currentThemeID = res.data;
          this.$request
            .get("/themes")
            .then((res) => {
              this.themes = res.data;
              for (let theme of this.themes) {
                if (theme.cover === '') {
                  theme.cover = require('@/assets/default.png')
                }
                if (theme.id === this.currentThemeID) {
                  theme.isCurrent = true;
                  this.currentTheme = theme;
                  this.themes.splice(this.themes.indexOf(theme), 1);
                  continue;
                }
              }
              this.themes.unshift(this.currentTheme);
            })
            .catch(() => {
              this.$message.error("获取主题列表失败");
            });
        })
        .catch(() => {
          this.$message.error("获取当前主题失败");
        });
    },
  },
  beforeMount() {
    this.fetchData();
  },
  components: {
    PageView,
    ThemeInstallDialog,
    SettingIcon,
    LockOnIcon,
    LockOffIcon,
    DeleteIcon,
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

.themes-container {
  display: flex;
  .card-container {
    margin: 0 10px;
    width: 400px;
  }
}
</style>