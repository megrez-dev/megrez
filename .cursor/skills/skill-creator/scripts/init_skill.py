#!/usr/bin/env python3
"""
技能初始化器 - 从模板创建新技能

用法:
    init_skill.py <skill-name> --path <path>

示例:
    init_skill.py my-new-skill --path skills/public
    init_skill.py my-api-helper --path skills/private
    init_skill.py custom-skill --path /custom/location
"""

import sys
from pathlib import Path
from datetime import datetime


SKILL_TEMPLATE = """---
name: {skill_name}
description: [TODO: 完整且信息丰富的技能描述，说明技能的功能和使用场景。包含何时使用此技能的具体场景、文件类型或触发任务。]
---

# {skill_title}

## 概述

[TODO: 1-2 句话解释此技能的功能]

## 技能结构选择

[TODO: 选择最适合此技能目的的结构。常见模式：

**1. 工作流程型** (适用于顺序流程)
- 适用于有明确步骤的流程
- 示例：DOCX 技能的"工作流程决策树" → "读取" → "创建" → "编辑"
- 结构：## 概述 → ## 工作流程决策树 → ## 步骤 1 → ## 步骤 2...

**2. 任务型** (适用于工具集合)
- 适用于技能提供不同操作/功能的情况
- 示例：PDF 技能的"快速开始" → "合并 PDF" → "拆分 PDF" → "提取文本"
- 结构：## 概述 → ## 快速开始 → ## 任务类别 1 → ## 任务类别 2...

**3. 参考/指南型** (适用于标准或规范)
- 适用于品牌指南、编码标准或要求
- 示例：品牌样式的"品牌指南" → "颜色" → "字体" → "特性"
- 结构：## 概述 → ## 指南 → ## 规范 → ## 用法...

**4. 能力型** (适用于集成系统)
- 适用于技能提供多个相互关联功能的情况
- 示例：产品管理的"核心能力" → 编号能力列表
- 结构：## 概述 → ## 核心能力 → ### 1. 功能 → ### 2. 功能...

可以根据需要混合和匹配模式。大多数技能会组合模式（例如，从任务型开始，为复杂操作添加工作流程）。

完成后删除整个"技能结构选择"部分 - 这只是指导。]

## [TODO: 根据选择的结构替换为第一个主要部分]

[TODO: 在此添加内容。参见现有技能中的示例：
- 技术技能的代码示例
- 复杂工作流程的决策树
- 带有实际用户请求的具体示例
- 根据需要引用 scripts/templates/references]

## 资源

此技能包含示例资源目录，演示如何组织不同类型的捆绑资源：

### scripts/
可直接运行以执行特定操作的可执行代码（Python/Bash 等）。

**适用于：** Python 脚本、shell 脚本或执行自动化、数据处理或特定操作的任何可执行代码。

**注意：** 脚本可以在不加载到上下文的情况下执行，但 Claude 仍可以读取它们以进行修补或环境调整。

### references/
旨在加载到上下文中以指导 Claude 流程和思考的文档和参考材料。

**适用于：** 深入文档、API 参考、数据库模式、综合指南或 Claude 在工作时应参考的任何详细信息。

### assets/
不打算加载到上下文中，而是在 Claude 产生的输出中使用的文件。

**适用于：** 模板、样板代码、文档模板、图像、图标、字体或任何用于复制或在最终输出中使用的文件。

---

**可以删除任何不需要的目录。** 并非每个技能都需要所有三种类型的资源。
"""

README_TEMPLATE = """## 技能文档

### 基本信息
- 技能名: `{skill_name}`
- 创建人: @[TODO: 填写创建者名称]
- 版本: v1.0.0
- 更新时间: {date}

### 适用场景
[TODO: 描述此技能适用的具体场景和用例]

### 前置条件
- [TODO: 列出使用此技能需要的前置条件]
- 例如：已配置某 MCP、已安装某依赖等

### 使用示例
```
[TODO: 示例触发语句或命令]
```

### 注意事项
⚠️ [TODO: 列出使用此技能时需要注意的事项]

### 已知问题
- [ ] [TODO: 待解决的问题]

### 相关技能
- [TODO: 列出相关技能及简要描述]
"""

EXAMPLE_SCRIPT = '''#!/usr/bin/env python3
"""
{skill_name} 的示例辅助脚本

这是一个可以直接执行的占位符脚本。
根据实际需要替换实现或删除此文件。
"""

def main():
    print("这是 {skill_name} 的示例脚本")
    # TODO: 在此添加实际脚本逻辑
    # 可以是数据处理、文件转换、API 调用等

if __name__ == "__main__":
    main()
'''

EXAMPLE_REFERENCE = """# {skill_title} 参考文档

这是详细参考文档的占位符。
根据实际需要替换内容或删除此文件。

## 何时使用参考文档

参考文档适用于：
- 全面的 API 文档
- 详细的工作流程指南
- 复杂的多步骤流程
- 对于主 SKILL.md 来说过长的信息
- 仅在特定用例中需要的内容

## 结构建议

### API 参考示例
- 概述
- 认证
- 带示例的端点
- 错误代码
- 速率限制

### 工作流程指南示例
- 前置条件
- 分步说明
- 常见模式
- 故障排除
- 最佳实践
"""

EXAMPLE_ASSET = """# 示例资产文件

此占位符表示资产文件的存储位置。
根据实际需要替换为实际资产文件（模板、图像、字体等）或删除此文件。

资产文件不打算加载到上下文中，而是在 Claude 产生的输出中使用。

## 常见资产类型

- 模板：.pptx、.docx、样板目录
- 图像：.png、.jpg、.svg、.gif
- 字体：.ttf、.otf、.woff、.woff2
- 样板代码：项目目录、启动文件
- 图标：.ico、.svg
- 数据文件：.csv、.json、.xml、.yaml

注意：这是文本占位符。实际资产可以是任何文件类型。
"""


def title_case_skill_name(skill_name):
    """将连字符分隔的技能名称转换为标题格式以供显示。"""
    return ' '.join(word.capitalize() for word in skill_name.split('-'))


def init_skill(skill_name, path):
    """
    初始化新的技能目录，包含模板 SKILL.md 和 README.md。

    参数:
        skill_name: 技能名称
        path: 技能目录应创建的路径

    返回:
        创建的技能目录路径，如果出错则返回 None
    """
    # 确定技能目录路径
    skill_dir = Path(path).resolve() / skill_name

    # 检查目录是否已存在
    if skill_dir.exists():
        print(f"❌ 错误: 技能目录已存在: {skill_dir}")
        return None

    # 创建技能目录
    try:
        skill_dir.mkdir(parents=True, exist_ok=False)
        print(f"✅ 已创建技能目录: {skill_dir}")
    except Exception as e:
        print(f"❌ 创建目录时出错: {e}")
        return None

    # 从模板创建 SKILL.md
    skill_title = title_case_skill_name(skill_name)
    skill_content = SKILL_TEMPLATE.format(
        skill_name=skill_name,
        skill_title=skill_title
    )

    skill_md_path = skill_dir / 'SKILL.md'
    try:
        skill_md_path.write_text(skill_content)
        print("✅ 已创建 SKILL.md")
    except Exception as e:
        print(f"❌ 创建 SKILL.md 时出错: {e}")
        return None

    # 创建 README.md
    readme_content = README_TEMPLATE.format(
        skill_name=skill_name,
        date=datetime.now().strftime("%Y-%m-%d")
    )

    readme_path = skill_dir / 'README.md'
    try:
        readme_path.write_text(readme_content)
        print("✅ 已创建 README.md")
    except Exception as e:
        print(f"❌ 创建 README.md 时出错: {e}")
        return None

    # 创建带有示例文件的资源目录
    try:
        # 创建 scripts/ 目录和示例脚本
        scripts_dir = skill_dir / 'scripts'
        scripts_dir.mkdir(exist_ok=True)
        example_script = scripts_dir / 'example.py'
        example_script.write_text(EXAMPLE_SCRIPT.format(skill_name=skill_name))
        example_script.chmod(0o755)
        print("✅ 已创建 scripts/example.py")

        # 创建 references/ 目录和示例参考文档
        references_dir = skill_dir / 'references'
        references_dir.mkdir(exist_ok=True)
        example_reference = references_dir / 'api_reference.md'
        example_reference.write_text(EXAMPLE_REFERENCE.format(skill_title=skill_title))
        print("✅ 已创建 references/api_reference.md")

        # 创建 assets/ 目录和示例资产占位符
        assets_dir = skill_dir / 'assets'
        assets_dir.mkdir(exist_ok=True)
        example_asset = assets_dir / 'example_asset.txt'
        example_asset.write_text(EXAMPLE_ASSET)
        print("✅ 已创建 assets/example_asset.txt")
    except Exception as e:
        print(f"❌ 创建资源目录时出错: {e}")
        return None

    # 打印后续步骤
    print(f"\n✅ 技能 '{skill_name}' 已成功初始化于 {skill_dir}")
    print("\n后续步骤:")
    print("1. 编辑 SKILL.md 完成 TODO 项并更新描述（使用中文）")
    print("2. 编辑 README.md 完成维护文档（使用中文）")
    print("3. 自定义或删除 scripts/、references/ 和 assets/ 中的示例文件")
    print("4. 准备好后运行验证器检查技能结构")

    return skill_dir


def main():
    if len(sys.argv) < 4 or sys.argv[2] != '--path':
        print("用法: init_skill.py <skill-name> --path <path>")
        print("\n技能名称要求:")
        print("  - 连字符分隔的标识符 (例如 'data-analyzer')")
        print("  - 仅限小写字母、数字和连字符")
        print("  - 最多 40 个字符")
        print("  - 必须与目录名完全匹配")
        print("\n示例:")
        print("  init_skill.py my-new-skill --path skills/public")
        print("  init_skill.py my-api-helper --path skills/private")
        print("  init_skill.py custom-skill --path /custom/location")
        sys.exit(1)

    skill_name = sys.argv[1]
    path = sys.argv[3]

    print(f"🚀 正在初始化技能: {skill_name}")
    print(f"   位置: {path}")
    print()

    result = init_skill(skill_name, path)

    if result:
        sys.exit(0)
    else:
        sys.exit(1)


if __name__ == "__main__":
    main()
