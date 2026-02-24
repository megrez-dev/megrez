# Megrez-Image组件

`版本: 1.0.0`

图片容器，可以默认预览，自定义自定义遮罩，自定义加载失败后兜底图片。

### Attributes

| 参数          | 说明                                                         | 类型             | 可选值                                     | 默认值 |
| ------------- | ------------------------------------------------------------ | ---------------- | ------------------------------------------ | ------ |
| src           | 同原生HTML，设置图片地址                                     | String           | ——                                         | -      |
| fit           | 确定图片如何适应容器框，同原生 [object-fit](https://developer.mozilla.org/en-US/docs/Web/CSS/object-fit) | String           | fill / contain / cover / none / scale-down | -      |
| fallback      | 配置兜底图片地址                                             | String           | ——                                         | -      |
| preview       | 配置是否可以预览，若设置为false则不使用默认预览，若设置为字符串类型，则自定义预览文案 | Boolean\|String  | true/false/any string                      | true   |
| viewerOptions | 配置v-viewer选项[viewerjs.option](https://github.com/fengyuanchen/viewerjs#options) | Object           | ——                                         | -      |
| maskClass     | 设置遮罩层css类名                                            | String\|String[] | ——                                         | -      |

### slot

|插槽名称|说明|

### slot

| 插槽名    | 说明                                                         |
| --------- | ------------------------------------------------------------ |
| maskInner | 用于自定义遮罩层内容，当设置了该slot之后，相当于将preview设置为false，可以通过maskClick函数监听元素被点击 |

### events

| 事件名称 | 说明             | 参数回调   |
| -------- | ---------------- | ---------- |
| load     | 图片加载成功触发 | (e: Event) |
| error    | 图片加载失败触发 | (e: Error) |

