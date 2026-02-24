<template>
  <div class="page-view">
    <div class="page-header">
      <BreadCrumb></BreadCrumb>
      <slot name="header"></slot>
    </div>
    <t-content class="content-wrapper">
      <slot name="content" />
    </t-content>
  </div>
</template>
<script>
import BreadCrumb from "@/layouts/BreadCrumb.vue";
export default {
  name: "PageView",
  components: { BreadCrumb },
};
</script>
<style lang="less" scoped>
.page-view {
  height: 100%;
}
.content-wrapper {
  padding: 24px;
}
.page-header {
  background-color: @bg-color-container;
  padding: 0 24px;
  display: flex;
  justify-content: space-between;
  flex-direction: column;
  .page-title {
    font-size: 18px;
    font-weight: 500;
    color: #333;
  }
  .page-actions {
    display: flex;
    align-items: center;
    .action-item {
      margin-left: 16px;
      color: #333;
      font-size: 14px;
      font-weight: 500;
      cursor: pointer;
      transition: color 0.2s cubic-bezier(0.38, 0, 0.24, 1);
      &:hover {
        color: #1890ff;
      }
    }
  }
}
</style>