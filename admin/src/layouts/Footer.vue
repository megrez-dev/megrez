<template>
  <div :class="prefix + '-footer'">Copyright @ 2021-{{ new Date().getFullYear() }} Megrez | Designed by AlkaidChen</div>
</template>

<script>
import Vue from 'vue';

export default Vue.extend({
  name: 'Footer',
  data() {
    return {
      prefix:'megrez',
    };
  },
});
</script>
<style lang="less" scoped>
@import '@/style/variables';

.megrez-footer {
  color: @text-color-placeholder;
  line-height: 20px;
  text-align: center;
}

</style>
