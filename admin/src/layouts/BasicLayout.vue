<template>
  <t-layout>
    <t-aside width="auto">
      <Aside></Aside>
    </t-aside>
    <t-layout>
      <t-header>
        <Header></Header>
      </t-header>
      <t-content>
        <Content></Content>
      </t-content>
    </t-layout>
  </t-layout>
</template>

<script>
import Aside from "@/layouts/Aside.vue";
import Header from "@/layouts/Header.vue";
import Content from "@/layouts/Content.vue";
import "@/style/layout.less";

export default {
  name: "BasicLayout",
  components: { Aside, Header, Content },
};
</script>
<style lang="less" scoped>
</style>