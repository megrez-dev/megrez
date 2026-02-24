<template>
  <div class="breadcrumb-container">
    <t-breadcrumb :maxItemWidth="'150'">
      <t-breadcrumbItem
        v-for="item in this.breadcrumbItems"
        :key="item.name"
        :to="{ name: item.name }"
      >
        <icon :name="icon" v-if="item.name === 'Index'" />
        {{ item.meta.title }}
      </t-breadcrumbItem>
    </t-breadcrumb>
  </div>
</template>

<script>
import { Icon } from 'tdesign-icons-vue';
export default {
  name: 'BreadCrumb',
  data() {
    return {
      icon: 'home',
    };
  },
  computed: {
    breadcrumbItems() {
      // 由于路由结构原因，需要手动添加面包屑的第一项
      let items = [
        {
          path: '/',
          name: 'Index',
          redirect: '/dashboard',
          meta: { title: '首页', icon: 'home' },
        },
      ];
      this.$route.matched.forEach((item) => {
        // 由于路由结构原因，如果是单级路由，如dashboard，则不需要将其子路由添加到面包屑中，即子路由中无meta.title
        if (!item.meta.title) return;
        if (item.meta.icon) this.icon = item.meta.icon;
        // 由于单级路由本身不设置那么属性，需要从子路由中读取
        if (item.single) item.name = item.children[0].name;
        items.push(item);
      });
      return items;
    },
  },
  components: {
    Icon,
  },
};
</script>
<style lang="less" scoped>
.breadcrumb-container {
  display: flex;
  align-items: center;
  height: 48px;
  .breadcrumb-link {
    text-decoration: none;
  }
}
</style>
