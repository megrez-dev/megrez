<template>
  <t-layout class="megrez-layout">
    <transition name="fade" mode="out-in">
      <router-view />
    </transition>
    <t-footer class="megrez-footer-layout">
      <Footer></Footer>
    </t-footer>
  </t-layout>
</template>

<script>
import Footer from "@/layouts/Footer.vue";
export default {
  name: "Content",
  components: { Footer },
};
</script>

<style lang="less" scoped>
.fade-leave-active,
.fade-enter-active {
  transition: opacity @anim-duration-slow @anim-time-fn-easing;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>