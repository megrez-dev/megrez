<template>
  <PageView>
    <template slot="content">
      <div class="links-container">
        <t-row>
          <t-col :flex="'39%'">
            <div class="card-container">
              <div class="card-title">
                <h2>添加友链</h2>
              </div>
              <div class="card-content">
                <t-form
                  :data="formData"
                  ref="form"
                  @reset="onReset"
                  @submit="onSubmit"
                  labelAlign="top"
                  :colon="true"
                >
                  <t-form-item label="姓名" name="name">
                    <t-input
                      v-model="formData.name"
                      placeholder="请输入内容"
                    ></t-input>
                  </t-form-item>
                  <t-form-item label="手机号码" name="tel">
                    <t-input
                      v-model="formData.tel"
                      placeholder="请输入内容"
                    ></t-input>
                  </t-form-item>
                  <t-form-item label="接收短信" name="status">
                    <t-switch v-model="formData.status"></t-switch>
                  </t-form-item>
                  <t-form-item label="性别" name="gender">
                    <t-radio-group v-model="formData.gender">
                      <t-radio value="1">男</t-radio>
                      <t-radio value="2">女</t-radio>
                    </t-radio-group>
                  </t-form-item>
                  <t-form-item label="课程" name="course">
                    <t-checkbox-group
                      v-model="formData.course"
                      :options="courseOptions"
                    ></t-checkbox-group>
                  </t-form-item>
                  <t-form-item>
                    <t-button
                      theme="primary"
                      type="submit"
                      style="margin-right: 10px"
                      >提交</t-button
                    >
                    <t-button theme="default" variant="base" type="reset"
                      >重置</t-button
                    >
                  </t-form-item>
                </t-form>
              </div>
            </div>
          </t-col>
          <t-col :flex="'1%'"> </t-col>
          <t-col :flex="'60%'">
            <div class="card-container">
              <div class="card-title">
                <h2>所有友链</h2>
              </div>
              <div class="card-content">
                <t-table
                  :data="data"
                  :columns="columns"
                  :rowKey="rowKey"
                  :verticalAlign="verticalAlign"
                  :loading="isLoading"
                  :pagination="pagination"
                  @change="rehandleChange"
                >
                  <template #op="slotProps">
                    <a class="t-button-link" @click="handleClickDetail()"
                      >编辑</a
                    >
                    <t-divider layout="vertical" />
                    <a
                      class="t-button-link"
                      @click="handleClickRecyle(slotProps)"
                      >回收站</a
                    >
                  </template>
                </t-table>
              </div>
            </div>
          </t-col>
        </t-row>
      </div>
    </template>
  </PageView>
</template>

<script>
import PageView from "@/layouts/PageView";
export default {
  name: "AttachmentList",
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
    }
  }
}
</style>