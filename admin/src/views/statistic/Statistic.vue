<template>
  <PageView>
    <template slot="content">
      <div class="links-container">
        <div class="card-container">
          <div class="card-title">
            <h2>统计</h2>
          </div>
          <div class="card-content">
            <h3>未实现</h3>
          </div>
        </div>
      </div>
    </template>
  </PageView>
</template>

<script>
import PageView from "@/layouts/PageView";
export default {
  name: "Statistic",
  data() {
    return {};
  },
  components: {
    PageView,
  },
};
</script>

<style lang="less" scoped>
.links-container {
  .card-container {
    padding: 0 0 24px 0;
    margin-right: 10px;
    background-color: @bg-color-container;
    border-radius: 2px;
    width: 100%;
    display: flex;
    flex-direction: column;
    .card-title {
      padding: 0 24px;
      border-bottom: 1px solid @border-level-1-color;
    }
    .card-content {
      padding: 16px 24px 0;
      height: 400px;
    }
  }
}
</style>