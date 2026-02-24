<template>
  <PageView>
    <template slot="content">
      <t-row>
        <t-col :flex="'34%'">
          <t-card title="添加友链" :bordered="false" header-bordered>
            <t-form ref="form" labelAlign="top" :colon="true">
              <t-form-item label="网站名称" name="name">
                <t-input v-model="link.name"></t-input>
              </t-form-item>
              <t-form-item label="网站地址" name="url">
                <t-input v-model="link.url" placeholder="http://"></t-input>
              </t-form-item>
              <t-form-item label="Logo" name="logo">
                <t-input v-model="link.logo"></t-input>
              </t-form-item>
              <t-form-item label="排序" name="priority">
                <t-input-number
                  v-model="link.priority"
                  theme="row"
                ></t-input-number>
              </t-form-item>
              <t-form-item label="网站描述" name="description">
                <t-textarea
                  v-model="link.description"
                  :autosize="{ minRows: 3 }"
                />
              </t-form-item>
            </t-form>
            <template slot="footer">
              <t-button
                theme="primary"
                @click="handleClickAdd"
                v-if="mode === 'add'"
              >
                添加
              </t-button>
              <t-button
                theme="primary"
                @click="handleClickUpdate"
                v-if="mode === 'edit'"
              >
                更新
              </t-button>
              <t-button
                theme="primary"
                variant="dashed"
                @click="handleClickReturn"
                v-if="mode === 'edit'"
              >
                返回添加
              </t-button>
            </template>
          </t-card>
        </t-col>
        <t-col :flex="'1%'"> </t-col>
        <t-col :flex="'65%'">
          <t-card title="友链列表" :bordered="false" header-bordered>
            <t-table
              :data="links"
              :columns="columns"
              rowKey="id"
              verticalAlign="middle"
              :loading="isLoading"
              :pagination="pagination"
              @change="rehandleChange"
            >
              <template #url="{ row }">
                <a
                  class="t-button-link"
                  :href="row.url"
                  target="_blank"
                  style="text-overflow: ellipsis"
                  >{{ row.url }}</a
                >
              </template>
              <template #logo="{ row }">
                <div class="logo-wrapper">
                  <img :src="row.logo" :alt="row.name" />
                </div>
              </template>
              <template #priority="{ row }">
                <t-badge
                  :count="row.priority"
                  shape="round"
                  :offset="[-14, -5]"
                  showZero
                >
                </t-badge>
              </template>
              <template #op="slotProps">
                <a class="t-button-link" @click="handleClickEdit(slotProps)"
                  >编辑</a
                >
                <t-divider layout="vertical" />
                <a class="t-button-link" @click="handleClickDelete(slotProps)"
                  >删除</a
                >
              </template>
            </t-table>
          </t-card>
        </t-col>
      </t-row>
    </template>
  </PageView>
</template>

<script>
import PageView from "@/layouts/PageView";
export default {
  name: "Links",
  data() {
    return {
      mode: "add",
      link: {
        name: "",
        url: "",
        logo: "",
        priority: 0,
        description: "",
      },
      links: [],
      isLoading: false,
      columns: [
        {
          colKey: "name",
          title: "名称",
          width: "150px",
        },
        {
          colKey: "url",
          title: "URL",
        },
        {
          colKey: "logo",
          title: "LOGO",
          width: "100px",
        },
        {
          colKey: "priority",
          title: "排序",
          width: "100px",
        },
        {
          fixed: "right",
          colKey: "op",
          title: "操作",
        },
      ],
      tableLayout: "auto",
      rowClassName: "property-class",
      pagination: {
        current: 1,
        pageSize: 5,
      },
    };
  },
  async mounted() {
    await this.fetchData(this.pagination);
  },

  methods: {
    handleClickAdd() {
      //validate
      if (this.link.name === "") {
        this.$message.warning("名称不能为空");
        return;
      }
      if (this.link.url === "") {
        this.$message.warning("URL不能为空");
        return;
      }
      if (this.link.logo === "") {
        this.$message.warning("LOGO不能为空");
        return;
      }
      this.$request
        .post("link", this.link)
        .then((res) => {
          this.$message.success("添加成功");
          this.links.push(res.data);
          this.clearForm();
        })
        .catch(() => {
          this.$message.warning("添加失败");
        });
    },
    async fetchData(pagination = this.pagination) {
      try {
        this.isLoading = true;
        this.$request
          .get("links", {
            params: {
              pageNum: pagination.current,
              pageSize: pagination.pageSize,
            },
          })
          .then((res) => {
            this.isLoading = false;
            this.links = res.data.list;
            this.pagination = {
              ...pagination,
              total: res.data.total,
            };
          });
      } catch (err) {
        this.links = [];
      }
    },
    // 也可以使用 page-change 事件
    async rehandleChange(changeParams) {
      const { current, pageSize } = changeParams.pagination;
      const pagination = { current, pageSize };
      await this.fetchData(pagination);
    },
    handleClickEdit({ row }) {
      this.mode = "edit";
      this.link = { ...row };
    },
    handleClickDelete({ row }) {
      this.$request
        .delete("link/" + row.id)
        .then(() => {
          this.$message.info("删除成功");
          for (let i = 0; i < this.links.length; i++) {
            if (this.links[i].id === row.id) {
              this.links.splice(i, 1);
              break;
            }
          }
        })
        .catch(() => {
          this.$message.error("删除失败");
        });
    },
    handleClickUpdate() {
      //validate
      if (this.link.name === "") {
        this.$message.warning("名称不能为空");
        return;
      }
      if (this.link.url === "") {
        this.$message.warning("URL不能为空");
        return;
      }
      if (this.link.logo === "") {
        this.$message.warning("LOGO不能为空");
        return;
      }
      this.$request
        .put("link/" + this.link.id, this.link)
        .then(() => {
          this.$message.info("更新成功");
          this.mode = "add";
          this.clearForm();
          this.fetchData();
        })
        .catch(() => {
          this.$message.error("更新失败");
        });
    },
    handleClickReturn() {
      this.mode = "add";
      this.clearForm();
    },
    clearForm() {
      this.link.id = 0;
      this.link.name = "";
      this.link.url = "";
      this.link.logo = "";
      this.link.priority = 0;
      this.link.description = "";
    },
  },
  components: {
    PageView,
  },
};
</script>

<style lang="less" scoped>
.logo-wrapper {
  width: 40px;
  height: 40px;
  background: #eeeeee;
  border-radius: 2px;
  img {
    margin: auto;
    width: 100%;
    height: 100%;
    border-radius: 2px;
  }
}
</style>