<template>
  <t-menu :defaultExpanded="expanded" :value="active" :theme="theme" :collapsed="collapsed">
    <template #logo>
      <div class="aside-title-container">
        <span class="aside-title-logo">M</span>
      </div>
    </template>
    <div class="menu-root-item" v-for="rootItem in manuList" :key="rootItem.name">
      <!-- 慎改rootItem.children[0].name！详情见/src/router/modules/dashboard.js注释 -->
      <t-menu-item
        :value="rootItem.path"
        :to="{ name: rootItem.name || rootItem.children[0].name }"
        v-if="!rootItem.children || !rootItem.children.length || rootItem.meta.single"
      >
        <template #icon>
          <icon :name="rootItem.meta.icon || rootItem.children[0].meta.icon" />
        </template>
        {{ rootItem.meta.title }}
      </t-menu-item>
      <t-submenu :value="rootItem.path" v-else>
        <template #icon>
          <icon :name="rootItem.meta.icon" />
        </template>
        <span slot="title">{{ rootItem.meta.title }}</span>
        <div class="menu-sub-item" v-for="subItem in rootItem.children" :key="subItem.name">
          <t-menu-item :value="subItem.path" :to="{ name: subItem.name }">
            <template #icon>
              <icon :name="subItem.meta.icon" />
            </template>
            <span>{{ subItem.meta.title }}</span>
          </t-menu-item>
        </div>
      </t-submenu>
    </div>
  </t-menu>
</template>

<script>
import { Icon } from 'tdesign-icons-vue';
import { menuRoutes } from '@/router/index.js';

export default {
  components: { Icon },
  data() {
    return {
      maxLevel: 3,
    };
  },
  computed: {
    active() {
      if (!this.$route.path) return '';
      const active = this.$route.path
        .split('/')
        .filter((_item, index) => index <= this.maxLevel && index > 0 && _item)
        .map((item) => `/${item}`)
        .join('');
      return active;
    },
    manuList() {
      return this.getMenuList(menuRoutes);
    },
    expanded() {
      return [`/${this.active.split('/')[1]}`];
    },
    collapsed() {
      return this.$store.state.app.collapsed;
    },
    theme() {
      if (this.$store.state.app.isDark) {
        return 'dark';
      } else {
        return 'light';
      }
    },
  },
  methods: {
    // 递归拼接path，方便根据active寻找正确的高亮项
    getMenuList(list, basePath) {
      if (!list) {
        return [];
      }
      return list.map((item) => {
        let path = '';
        // 如果以'/'或'http'开头，认为其path为绝对路径，不需要与父级路径拼接，适应'/article/list'这种写法，但不推荐
        if (!path.startsWith('/') || !path.startsWith('http')) {
          path = basePath ? `${basePath}/${item.path}` : item.path;
        } else {
          path = item.path;
        }
        return {
          ...item,
          path,
          children: this.getMenuList(item.children, path),
        };
      });
    },
  },
};
</script>
<style lang="less">
aside {
  position: fixed;
  top: 0;
  z-index: 1;
  box-shadow: 2px 0 8px 0 rgb(29 35 41 / 5%);
}
.aside-title-container {
  text-align: center;
  flex: 1;
}
.aside-title-logo {
  font-size: 35px;
  font-weight: 1000;
  color: #0052d9;
}
.aside-title-text {
  font-size: 30px;
  font-weight: 500;
}
</style>
