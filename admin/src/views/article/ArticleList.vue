<template>
  <PageView>
    <template slot="content">
      <t-button @click="handleWrite">
        <add-icon slot="icon" />
        写文章
      </t-button>
      <t-table
        class="article-list"
        rowKey="id"
        verticalAlign="middle"
        :data="articleList"
        :columns="columns"
        :loading="isArticleListLoading"
        :pagination="pagination"
        @change="rehandleChange"
      >
        <template #title="{ row }">
          <t-tooltip
            class="article-title-tooltips"
            theme="light"
            placement="mouse"
          >
            <span slot="content"> {{ row.title }} <span style="color: rgba(0, 0, 0, .4)">更新于{{ row.editTime | timeAgo }}</span></span>
            <a class="t-button-link text-ellipsis" :href="'/article/' + row.id" target="_blank">
              {{ row.title }}
            </a>
          </t-tooltip>
        </template>
        <template #status="{ row: { status } }">
          <t-badge dot :offset="['120%', '50%']" :color="status | statusColor">
            {{ status | statusText }}
          </t-badge>
        </template>
        <template #categories="{ row }">
          <t-tag
            v-for="(category, index) in row.categories"
            :key="index"
            size="small"
            theme="primary"
            variant="light"
            style="margin-bottom: 8px; margin-right: 8px"
          >
            {{ category.name }}
          </t-tag>
        </template>
        <template #tags="{ row }">
          <t-tag
            v-for="(tag, index) in row.tags"
            :key="index"
            size="small"
            theme="success"
            variant="light"
            style="margin-bottom: 8px; margin-right: 8px"
          >
            {{ tag.name }}
          </t-tag>
        </template>
        <template #publishTime="{ row }">
          {{ row.publishTime | timeAgo }}
        </template>
        <template #op="{ row: { id } }">
          <t-button variant="text" size="small" @click="handleClickDetail(id)">
            <EditIcon size="18px" />
          </t-button>
          <t-button variant="text" size="small" @click="handleSetting(id)">
            <SettingIcon size="18px" />
          </t-button>
          <t-popconfirm @confirm="handleClickDelete(id)" content="确认删除？" theme="danger">
            <t-button variant="text" theme="danger" size="small" :loading="deleting">
              <DeleteIcon size="18px" v-show="!deleting"/>
            </t-button>
          </t-popconfirm>
        </template>
      </t-table>
    </template>
  </PageView>
</template>

<script>
import { AddIcon, EditIcon, DeleteIcon, SettingIcon } from 'tdesign-icons-vue';
import PageView from '@/layouts/PageView';
import { ARTICLE_STATUS_MAP } from '@/views/article/constants';

export default {
  name: 'ArticleList',
  components: { AddIcon, EditIcon, SettingIcon, DeleteIcon, PageView },
  data() {
    return {
      articleList: [],
      isArticleListLoading: false,
      deleting: false,
      columns: [
        { colKey: 'title', title: '标题', width: 220 },
        { colKey: 'status', title: '状态', width: 100 },
        { colKey: 'categories', title: '分类' },
        { colKey: 'tags', title: '标签' },
        { colKey: 'commentsNum', title: '评论', width: 80, align: 'center' },
        { colKey: 'visits', title: '访问', width: 80, align: 'center' },
        { colKey: 'publishTime', title: '发布时间', width: 180 },
        { colKey: 'op', title: '操作', width: 180, align: 'center', fixed: 'right' },
      ],
      pagination: {
        current: 1,
        pageSize: 10,
      },
    };
  },
  filters: {
    // 用于获取
    statusColor(status) {
      return ARTICLE_STATUS_MAP[status].color || 'red';
    },
    statusText(status) {
      return ARTICLE_STATUS_MAP[status].text || '未知';
    },
  },
  mounted() {
    this.listArticles(this.pagination);
  },
  methods: {
    listArticles(pagination = this.pagination) {
      this.isArticleListLoading = true;
      const { current, pageSize } = pagination;
      this.$request
        .get('articles?pageNum=' + current + '&pageSize=' + pageSize)
        .then((res) => {
          this.articleList = res.data.list ? res.data.list : [];
          this.pagination = {
            ...pagination,
            total: res.data.total,
          };
        })
        .catch(() => {
          this.$message.error('获取文章列表失败');
        })
        .finally(() => {
          this.isArticleListLoading = false;
        });
    },
    rehandleChange(pageInfo) {
      this.pagination.current = pageInfo.pagination.current;
      this.pagination.pageSize = pageInfo.pagination.pageSize;
      this.listArticles(this.pagination);
    },
    handleWrite() {
      this.$router.push({ name: 'ArticleEdit' });
    },
    handleClickDetail(id) {
      this.$router.push({
        name: 'ArticleEdit',
        query: { articleID: id },
      });
    },
    handleClickDelete(id) {
      this.deleting = true;
      this.$request
        .delete('article/' + id)
        .then((res) => {
          if (res.status === 0) {
            for (var i = 0; i < this.articleList.length; i++) {
              if (this.articleList[i].id === id) {
                this.articleList.splice(i, 1);
                break;
              }
            }
            this.$message.success('删除成功');
          }
        })
        .catch(() => {
          this.$message.warning('删除失败');
        })
        .finally(() => {
          this.deleting = false;
        });
    },
    handleClickSetting(id) {
      return id;
    },
  },
};
</script>
<style lang="less" scoped>
.article-list {
  padding: 16px 24px;
  margin: 16px 0;
}
</style>
