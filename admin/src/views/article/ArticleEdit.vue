<template>
  <PageView>
    <t-row
      class="article-edit-bar"
      slot="header"
      justify="space-between"
      align="bottom"
      :gutter="[0, 10]"
    >
      <t-col flex="1 1 auto">
        <h1>{{ article.title === '' ? '新文章' : article.title }}</h1>
        <t-tag v-if="editMode === 1" :theme="article.status | statusTheme" variant="light">{{
          article.status | statusText
        }}</t-tag>
        <span v-if="editMode === 1" class="last-edit-time"
          >更新于 {{ article.editTime | timeAgo }}</span
        >
      </t-col>
      <t-col flex="1 0 auto">
        <t-row justify="end">
          <t-button
            :loading="saving"
            :disabled="publishing || !loaded"
            theme="danger"
            variant="base"
            @click="saveDraft"
          >
            {{ article.status != 1 ? '存为草稿' : '保存草稿' }}
          </t-button>
          <t-button
            theme="primary"
            variant="base"
            :loading="publishing"
            :disabled="saving || !loaded"
            @click="handlePublish"
          >
            {{ editMode === 1 && article.status === 0 ? '更新' : '发布' }}
          </t-button>
        </t-row>
      </t-col>
    </t-row>
    <div class="article-edit-container" slot="content">
      <div class="article-edit-title">
        <t-input size="large" placeholder="请输入标题" v-model="article.title" />
      </div>
      <div class="vditor-container">
        <Vditor 
          ref="vditor"
          v-model="article.originalContent"
          :loading="!loaded"
          @countWord="countWord"
          @insertImage="openAttachesSelectDrawer"
        />
      </div>
      <t-tabs defaultValue="basic">
        <t-tab-panel value="basic">
          <template #label> <icon name="setting" style="margin-right: 4px" /> 基本设置 </template>
          <t-form class="form-container" ref="form" labelAlign="top" colon>
            <t-row :gutter="[50, 20]" justify="center">
              <t-col :xl="4" :span="6">
                <t-form-item
                  label="文章别名"
                  name="slug"
                  :help="`${origin}/article/${article.slug || '{slug}'}`"
                >
                  <t-input
                    v-model="article.slug"
                    placeholder="不填写则会默认使用标题拼音或文章id"
                  ></t-input>
                </t-form-item>
                <t-form-item label="分类" name="categories">
                  <t-select
                    v-model="article.categories"
                    filterable
                    multiple
                    :minCollapsedNum="5"
                    :options="categoryOptions"
                  />
                  <t-button theme="primary" @click="triggerAddCategoryForm(true)">
                    新建分类
                  </t-button>
                </t-form-item>
                <t-dialog
                  :visible="showAddCategoryForm"
                  header="新建分类"
                  @close="triggerAddCategoryForm(false)"
                  @confirm="handleCreateCategory"
                >
                  <t-form :data="newCategory" labelAlign="top" colon>
                    <t-form-item label="分类名称" name="name">
                      <t-input v-model="newCategory.name" placeholder="分类名称"></t-input>
                    </t-form-item>
                    <t-form-item
                      label="分类别名"
                      name="slug"
                      :help="`${origin}/category/${newCategory.slug || '{slug}'}`"
                    >
                      <t-input v-model="newCategory.slug" placeholder="分类别名"></t-input>
                    </t-form-item>
                  </t-form>
                </t-dialog>
                <t-form-item label="标签" name="tags">
                  <t-select
                    v-model="article.tags"
                    creatable
                    filterable
                    multiple
                    :minCollapsedNum="3"
                    :options="tagOptions"
                    @create="handleCreateTag"
                  />
                </t-form-item>
                <t-form-item label="摘要" name="summary">
                  <t-textarea
                    v-model="article.summary"
                    placeholder="若不填写，将会从文章中自动截取"
                    name="summary"
                    :autosize="{ minRows: 6 }"
                  />
                </t-form-item>
              </t-col>
              <t-col :xl="4" :span="6">
                <!-- 添加padding使其对齐左边的输入框 -->
                <t-row style="padding: 32px 0 24px 0;">
                  <t-col :span="6">
                    <t-form-item label="开启评论" name="allowedComment" labelAlign="left">
                      <t-switch v-model="article.allowedComment"></t-switch>
                    </t-form-item>
                  </t-col>
                  <t-col :span="6">
                    <t-form-item label="是否置顶" name="isTop" labelAlign="left">
                      <t-switch v-model="article.isTop"></t-switch>
                    </t-form-item>
                  </t-col>
                </t-row>
                <t-form-item label="封面图" name="cover">
                  <t-input
                    clearable
                    block
                    v-model="article.cover"
                    placeholder="点击封面选择图片，或者输入外部链接"
                  ></t-input>
                </t-form-item>
                <div class="article-cover-img">
                  <m-image
                    @maskClick="openCoverSelectDrawer"
                    :src="coverUrl"
                    fit="cover"
                    :preview="false"
                  ></m-image>
                </div>
              </t-col>
            </t-row>
          </t-form>
        </t-tab-panel>
        <t-tab-panel value="advanced">
          <template #label> <icon name="internet" style="margin-right: 4px" /> 高级设置 </template>
          <t-form class="form-container" ref="form" labelAlign="top" :colon="true">
            <t-row :gutter="[50, 20]" justify="center">
              <t-col :xl="2" :span="4">
                <t-form-item label="访问密码" name="password">
                  <t-input type="password" v-model="article.password" placeholder=""> </t-input>
                </t-form-item>
              </t-col>
              <t-col :xl="5" :span="8">
                <t-form-item label="SEO 关键字" name="keywords">
                  <t-select
                    v-model="article.seoKeywords"
                    creatable
                    filterable
                    multiple
                    placeholder="若不填写，将使用标签作为关键字"
                    :options="seoKeywordOptions"
                    :minCollapsedNum="3"
                    @create="createSEOKeyword"
                  />
                </t-form-item>
              </t-col>
              <t-col :xl="7" :span="12">
                <t-form-item label="SEO 描述" name="description">
                  <t-textarea
                    v-model="article.seoDescription"
                    placeholder="若不填写，将使用摘要作为描述"
                    name="seoDescription"
                    :autosize="{ minRows: 6 }"
                  />
                </t-form-item>
              </t-col>
            </t-row>
          </t-form>
        </t-tab-panel>
      </t-tabs>
      <!-- 此处展示了AttachSelectDrawer展示与否的两种用法 -->
      <!-- 方法1：使用.sync修饰符，当Drawer被关闭时自动更新visible的值 -->
      <AttachSelectDrawer
        ref="attachesSelectDrawer"
        mode="multiple"
        :visible.sync="attachesDrawerVisible"
        @select="handleAttachSelect"
      ></AttachSelectDrawer>
      <!-- 方法2：使用close事件，在close事件中手动更改visible值 -->
      <AttachSelectDrawer
        ref="coverSelectDrawer"
        mode="single"
        :visible="coverDrawerVisible"
        @close="closeCoverSelectDrawer"
        @select="handleCoverSelect"
      ></AttachSelectDrawer>
    </div>
  </PageView>
</template>

<script>
import Vditor from '@/views/article/components/Vditor.vue';
import AttachSelectDrawer from '@/components/attachment/AttachSelectDrawer.vue';
import PageView from '@/layouts/PageView';
import MImage from '@/components/image/Image.vue';
import { Icon } from 'tdesign-icons-vue';
import { EDIT_MODE, DEFAULT_COVER_URL, ARTICLE_STATUS_MAP } from '@/views/article/constants';

export default {
  name: 'ArticleEdit',
  components: { Vditor, AttachSelectDrawer, Icon, PageView, MImage },
  data() {
    return {
      loaded: false,
      origin: location.origin,
      publishing: false, // 表示正在上传或正在更新
      saving: false, // 表示正在存草稿
      article: {
        title: '',
        originalContent: ' ',
        formatContent: '',
        slug: '',
        allowedComment: true,
        isTop: false,
        categories: [],
        tags: [],
        summary: '',
        cover: '',
        password: '',
        wordCount: 0,
        seoKeywords: [],
        seoDescription: '',
        status: 0,
        editTime: '',
      },
      preContent: ' ',
      preTitle: '',
      categoryOptions: [],
      tagOptions: [],
      seoKeywordOptions: [],
      showAddCategoryForm: false,
      newCategory: {
        name: '',
        slug: '',
      },
      attachesDrawerVisible: false,
      coverDrawerVisible: false,
    };
  },
  watch: {
    // 在文章编辑页跳转，只改变路由参数而不改变路由时，也需要更新默认的文章
    $route() {
      this.initData();
      this.setDefultArticle();
    },
  },
  computed: {
    articleID() {
      return this.$route.query.articleID;
    },
    editMode() {
      return this.articleID === undefined ? EDIT_MODE.create : EDIT_MODE.edit;
    },
    coverUrl() {
      return this.article.cover || DEFAULT_COVER_URL;
    },
  },
  methods: {
    initData() {
      this.loaded = false; // 标志着vditor加载完成
      this.publishing = false;
      this.saving = false;
      this.article = {
        title: '',
        originalContent: ' ', // 如果初始化时value为空或undefined，vditor会给一个默认值。
        formatContent: '',
        slug: '',
        allowedComment: true,
        isTop: false,
        categories: [],
        tags: [],
        summary: '',
        cover: '',
        password: '',
        wordCount: 0,
        seoKeywords: [],
        seoDescription: '',
        status: 0,
        editTime: '',
      };
      this.preContent = ' ';
      this.preTitle = '';
      this.categoryOptions = [];
      this.tagOptions = [];
      this.seoKeywordOptions = [];
      this.showAddCategoryForm = false;
      this.newCategory = {
        name: '',
        slug: '',
      };
    },
    // get article from server
    setDefultArticle() {
      if (!this.articleID) {
        this.$nextTick(() => {
          this.$refs.vditor.initViditor(() => {
            // 为了防止在未加载完成之前点击保存、发布，导致浏览器报错
            this.loaded = true;
          });
        });
        return;
      }
      // 根据articleId从后台获取初始化数据
      this.$request
        .get(`article/${this.articleID}`)
        .then((res) => {
          this.article = res.data;
          // 记录初始化的内容和标题。
          this.preContent = res.data.originalContent;
          this.preTitle = res.data.title;
          // 由于data的更新是异步的，所以需要在nextTick中执行，确保获取到最新的article
          this.$nextTick(() => {
            this.$refs.vditor.initViditor(() => {
              // 为了防止在未加载完成之前点击保存、发布，导致浏览器报错
              this.loaded = true;
            });
          });
        })
        .catch(() => {
          this.$message.warning('获取文章详情失败');
        });
    },
    // get word count from sub component Vditor
    countWord(length) {
      this.article.wordCount = length;
    },
    getFormatContent() {
      return this.$refs.vditor.contentEditor.getHTML();
    },
    saveDraft() {
      //validate
      if (this.article.title === '') {
        this.$message.warning('文章标题不能为空');
        return;
      }
      this.saving = true;
      this.article.formatContent = this.getFormatContent();
      this.article.status = 1;
      if (this.editMode === EDIT_MODE.create) {
        this.$request
          .post('article', this.article)
          .then((res) => {
            if (res.status === 0) {
              this.$message.success('保存成功');
              this.$router.push({ name: 'ArticleList', params: { hasSaved: true } });
            }
          })
          .catch(() => {
            this.$message.warning('保存失败');
          })
          .finally(() => {
            this.saving = false;
          });
      } else {
        this.$request
          .put('article/' + this.article.id, this.article)
          .then((res) => {
            if (res.status === 0) {
              this.$message.success('保存成功');
              this.$router.push({ name: 'ArticleList', params: { hasSaved: true } });
            }
          })
          .catch(() => {
            this.$message.warning('保存失败');
          })
          .finally(() => {
            this.saving = false;
          });
      }
    },
    handlePublish() {
      //validate
      if (this.article.title === '') {
        this.$message.warning('文章标题不能为空');
        return;
      }
      this.publishing = true;
      this.article.formatContent = this.getFormatContent();
      this.article.status = 0;
      // create
      if (this.editMode === EDIT_MODE.create) {
        this.$request
          .post('article', this.article)
          .then((res) => {
            if (res.status === 0) {
              this.$message.success('发布成功');
              this.$router.push({ name: 'ArticleList', params: { hasSaved: true } });
            }
          })
          .catch(() => {
            this.$message.warning('发布失败');
          })
          .finally(() => {
            this.publishing = false;
          });
      } else {
        // edit
        this.$request
          .put('article/' + this.article.id, this.article)
          .then((res) => {
            if (res.status === 0) {
              this.$message.success('更新成功');
              this.$router.push({ name: 'ArticleList', params: { hasSaved: true } });
            }
          })
          .catch(() => {
            this.$message.warning('更新失败');
          })
          .finally(() => {
            this.publishing = false;
          });
      }
    },
    triggerAddCategoryForm(ifShow = false) {
      this.showAddCategoryForm = ifShow;
    },
    openAttachesSelectDrawer() {
      this.attachesDrawerVisible = true;
    },
    closeAttachesSelectDrawer() {
      this.attachesDrawerVisible = false;
    },
    openCoverSelectDrawer() {
      this.coverDrawerVisible = true;
    },
    closeCoverSelectDrawer() {
      this.coverDrawerVisible = false;
    },
    handleCreateCategory() {
      // remove whitespace
      this.newCategory.name = this.newCategory.name.replace(/\s*/g, '');
      this.newCategory.slug = this.newCategory.slug.replace(/\s*/g, '');
      if (this.newCategory.name === '') {
        this.$message.warning('分类名称不能为空');
        return;
      }
      if (this.newCategory.slug === '') {
        this.$message.warning('分类别名不能为空');
        return;
      }
      this.$request
        .post('category', this.newCategory)
        .then((res) => {
          if (res.status === 0) {
            this.$message.success('创建分类成功');
            this.categoryOptions.unshift({
              value: res.data.id,
              label: res.data.name,
            });
            this.cancelCreateCategory();
          }
        });
    },
    cancelCreateCategory() {
      this.newCategory.name = '';
      this.newCategory.slug = '';
      this.showAddCategoryForm = false;
    },
    handleCreateTag(value) {
      const tagName = value.replace(/\s*/g, '');
      if (tagName === '') return this.$message.warning('标签不能为空');
      this.$request
        .post('tag', { name: tagName, slug: tagName })
        .then((res) => {
          this.$message.success('创建标签成功');
          this.tagOptions.push({
            value: res.data.id,
            label: res.data.name,
          });
          // 由于添加至article.tags中的应该是id，而不是value，value仅用于创建tag
          this.article.tags.splice(this.article.tags.indexOf(value), 1);
          this.article.tags.push(res.data.id);
        })
        .catch(() => {
          this.article.tags.splice(this.article.tags.indexOf(value), 1);
        });
    },
    createSEOKeyword(value) {
      this.seoKeywordOptions.push({
        value,
        label: value,
      });
    },
    handleAttachSelect(attaches) {
      const willInsertValue = attaches.map(attach => `![${attach.fileName}](${attach.url})`).join('\n');
      this.$refs.vditor.contentEditor.insertValue(willInsertValue);
    },
    handleCoverSelect(attach) {
      this.article.cover = attach.url;
    },
    // TODO: 还需要监听浏览器窗口的关闭和刷新
    handleRouteChange(to, from, next) {
      // 如果标题和内容均没有发生变化，则认为没有发生变化，不进行提示
      if (this.preContent === this.article.originalContent && this.preTitle === this.article.title)
        return next();
      // 如果发生变化，则判断params中的hasSaved，若hasSaved为true则说明是点击保存或点击发布按钮后引起的路由跳转。
      if (to.params.hasSaved) return next();
      const dialog = this.$dialog({
        header: '当前页面数据未保存，确定要离开吗？',
        body: '如果离开当前页面，您的数据将会丢失！',
        onConfirm() {
          dialog.hide();
          next();
        },
        theme: 'danger',
      });
    },
  },

  filters: {
    statusTheme(status) {
      return ARTICLE_STATUS_MAP[status].theme || 'danger';
    },
    statusText(status) {
      return ARTICLE_STATUS_MAP[status].text || '未知';
    },
  },

  mounted() {
    // list categories
    this.$request.get('categories').then((res) => {
      res.data.forEach((category) => {
        this.categoryOptions.push({
          value: category.id,
          label: category.name,
        });
      });
    });

    // list tags
    this.$request.get('tags').then((res) => {
      res.data.forEach((tag) => {
        this.tagOptions.push({
          value: tag.id,
          label: tag.name,
        });
      });
    });
    this.setDefultArticle();
  },

  beforeRouteUpdate(to, from, next) {
    this.handleRouteChange(to, from, next);
  },

  beforeRouteLeave(to, from, next) {
    this.handleRouteChange(to, from, next);
  },
};
</script>

<style lang="less" scoped>
.article-edit-bar {
  margin-bottom: 10px;
  flex-wrap: nowrap;
  h1 {
    margin-bottom: 8px;
  }
  .last-edit-time {
    margin-left: 10px;
    color: var(--td-text-color-placeholder);
  }
}

.article-edit-container {
  .article-edit-title {
    margin-bottom: 20px;
  }

  .vditor-container {
    background-color: var(--td-bg-color-container);
    margin-bottom: 20px;
  }
  .form-container {
    padding: 30px 30px;
  }
  .article-cover-img {
    width: 100%;
    height: 265px;
    cursor: pointer;
  }
}
</style>
