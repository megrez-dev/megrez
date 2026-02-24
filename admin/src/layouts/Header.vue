<template>
  <t-head-menu :theme="theme">
    <template #logo>
      <div class="header-operator-left">
        <t-button shape="square" variant="text" @click="toggleCollapsed">
          <t-icon size="20px" class="collapsed-icon" :name="collapsIcon" />
        </t-button>
      </div>
    </template>
    <template #operations>
      <div class="header-operator-right">
        <t-tooltip placement="bottom" content="代码仓库">
          <t-button shape="square" variant="text">
            <t-icon size="20px" name="logo-github" />
          </t-button>
        </t-tooltip>
        <t-tooltip placement="bottom" content="帮助文档">
          <t-button shape="square" variant="text">
            <t-icon size="20px" name="help-circle" />
          </t-button>
        </t-tooltip>
        <div class="theme-tab">
          <t-radio-group
            variant="default-filled"
            size="small"
            default-value="light"
            @change="toggleTheme"
          >
            <t-radio-button value="light"
              ><icon name="heart-filled" size="20px" style="color: orange"
            /></t-radio-button>
            <t-radio-button value="dark"
              ><icon name="star-filled" size="20px" style="color: gray"
            /></t-radio-button>
          </t-radio-group>
        </div>
      </div>
    </template>
  </t-head-menu>
</template>

<script>
import { Icon } from "tdesign-icons-vue";

export default {
  components: {
    Icon,
  },
  computed: {
    collapsIcon() {
      if (this.$store.state.app.collapsed) {
        return "menu-fold";
      } else {
        return "menu-unfold";
      }
    },
    theme() {
      if (this.$store.state.app.isDark) {
        return "dark";
      } else {
        return "light";
      }
    },
  },
  methods: {
    toggleCollapsed() {
      this.$store.commit("TOGGLE_COLLAPSED");
    },
    toggleTheme(value) {
      if (value === "light") {
        document.documentElement.removeAttribute("theme-mode");
      }
      if (value === "dark") {
        document.documentElement.setAttribute("theme-mode", "dark");
      }
      // TODO: 设置主题
      document.documentElement.setAttribute("theme-color", "default");
      this.$store.commit("TOGGLE_THEME");
    },
  },
};
</script>
<style lang="less">
header {
  box-shadow: 0 1px 4px rgb(0 21 41 / 8%);
  z-index: 1;
}

.header-operator-left {
  display: flex;
  margin-left: 20px;
  align-items: normal;
  line-height: 0;

  .collapsed-icon {
    font-size: 20px;
  }
}

.header-operator-right {
  display: flex;
  align-items: center;
  margin-right: 20px;

  .t-button {
    margin: 0 8px;
    &.header-user-btn {
      margin: 0;
    }
  }

  .theme-tab {
    display: flex;
    align-items: center;
    width: 60px;
    height: 32px;
    margin-left: 12px;

    .t-radio-group {
      height: 32px;
      .t-radio-button {
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 4px;
        height: 28px;
        width: 28px;
        .t-icon {
          font-size: 16px;
        }
      }
    }
  }
}

.t-head-menu__inner {
  border-bottom: 1px solid @border-level-1-color;
}

.t-menu--dark {
  .t-head-menu__inner {
    border-bottom: 1px solid var(--td-gray-color-10);
  }

  .t-button {
    --ripple-color: var(--td-gray-color-10) !important;
    &:hover {
      background: var(--td-gray-color-12) !important;
    }
  }
}
</style>
