---
name: skill-creator
description: 技能创建指南。当用户想要创建新技能（或更新现有技能）以扩展 Claude 的能力时使用此技能，包括专业知识、工作流程或工具集成。
license: Complete terms in LICENSE.txt
---

# 技能创建器

本技能提供创建高效技能的指导。

## 关于技能

技能是模块化、自包含的包，通过提供专业知识、工作流程和工具来扩展 Claude 的能力。可以将它们视为特定领域或任务的"入职指南"——它们将 Claude 从通用代理转变为配备程序性知识的专业代理。

### 技能提供的内容

1. **专业工作流程** - 特定领域的多步骤流程
2. **工具集成** - 处理特定文件格式或 API 的指令
3. **领域专业知识** - 公司特定知识、数据模式、业务逻辑
4. **捆绑资源** - 用于复杂和重复任务的脚本、参考资料和资产

## 核心原则

### 简洁为王

上下文窗口是公共资源。技能与 Claude 需要的所有其他内容共享上下文窗口：系统提示、对话历史、其他技能的元数据和实际用户请求。

**默认假设：Claude 已经非常聪明。** 只添加 Claude 还不具备的上下文。质疑每条信息："Claude 真的需要这个解释吗？"和"这段话值得消耗 token 吗？"

优先使用简洁的示例而非冗长的解释。

### 设置适当的自由度

根据任务的脆弱性和可变性匹配具体程度：

- **高自由度（文本指令）**：当多种方法都有效、决策取决于上下文或启发式方法指导时使用
- **中等自由度（伪代码或带参数的脚本）**：当存在首选模式、可接受一些变化或配置影响行为时使用
- **低自由度（特定脚本，少量参数）**：当操作脆弱易出错、一致性至关重要或必须遵循特定顺序时使用

### 技能结构

每个技能由必需的 SKILL.md 文件、必需的 README.md 文件和可选的捆绑资源组成：

```
skill-name/
├── SKILL.md (必需) - AI 代理使用的技能指令
├── README.md (必需) - 人类可读的维护文档
└── 捆绑资源 (可选)
    ├── scripts/          - 可执行代码 (Python/Bash 等)
    ├── references/       - 按需加载到上下文的文档
    └── assets/           - 输出中使用的文件 (模板、图标、字体等)
```

#### SKILL.md (必需)

每个 SKILL.md 包含：

- **Frontmatter** (YAML)：包含 `name` 和 `description` 字段。这是 Claude 判断何时使用技能的唯一字段，因此清晰全面地描述技能是什么以及何时使用非常重要。
- **正文** (Markdown)：使用技能的指令和指导。仅在技能触发后加载。

**重要**：SKILL.md 必须使用中文撰写，包括 frontmatter 中的 description 字段。

#### README.md (必需)

每个技能必须包含一个 README.md 文件，用于人类可读的维护文档。此文件不会被 AI 代理加载到上下文中，而是为技能维护者提供参考。

README.md 必须遵循以下模板格式：

```markdown
## 技能文档

### 基本信息
- 技能名: `skill-name`
- 创建人: @创建者名称
- 版本: v1.0.0
- 更新时间: YYYY-MM-DD

### 适用场景
描述此技能适用的具体场景和用例

### 前置条件
- 列出使用此技能需要的前置条件
- 例如：已配置某 MCP、已安装某依赖等

### 使用示例
\`\`\`
示例触发语句或命令
\`\`\`

### 注意事项
⚠️ 列出使用此技能时需要注意的事项

### 已知问题
- [ ] 待解决的问题
- [x] 已解决的问题 (版本号)

### 相关技能
- `related-skill-1`: 简要描述
- `related-skill-2`: 简要描述
```

#### 捆绑资源 (可选)

##### Scripts (`scripts/`)

用于需要确定性可靠性或重复编写的任务的可执行代码。

- **何时包含**：当相同代码被重复编写或需要确定性可靠性时
- **示例**：`scripts/rotate_pdf.py` 用于 PDF 旋转任务
- **优点**：Token 高效、确定性、可在不加载到上下文的情况下执行

##### References (`references/`)

按需加载到上下文中的文档和参考材料。

- **何时包含**：用于 Claude 在工作时应参考的文档
- **示例**：`references/finance.md` 用于财务模式，`references/api_docs.md` 用于 API 规范
- **优点**：保持 SKILL.md 精简，仅在 Claude 确定需要时加载

##### Assets (`assets/`)

不打算加载到上下文中，而是在 Claude 产生的输出中使用的文件。

- **何时包含**：当技能需要在最终输出中使用的文件时
- **示例**：`assets/logo.png` 用于品牌资产，`assets/template/` 用于模板文件

### 渐进式披露原则

技能使用三级加载系统来高效管理上下文：

1. **元数据 (name + description)** - 始终在上下文中 (~100 词)
2. **SKILL.md 正文** - 技能触发时 (<5k 词)
3. **捆绑资源** - Claude 按需加载

保持 SKILL.md 正文在 500 行以内以最小化上下文膨胀。接近此限制时将内容拆分到单独的文件中。

## 技能创建流程

技能创建包含以下步骤：

1. 通过具体示例理解技能
2. 规划可复用的技能内容 (scripts, references, assets)
3. 分析并确定分类目录
4. 初始化技能 (运行 init_skill.py)
5. 编辑技能 (实现资源并编写 SKILL.md 和 README.md)
6. 更新映射表 (SKILLS_COMMAND_MAP.md)
7. 基于实际使用迭代

按顺序执行这些步骤，仅在有明确理由时跳过。

### 步骤 1：通过具体示例理解技能

仅当技能的使用模式已经清楚理解时才跳过此步骤。

要创建有效的技能，需要清楚理解技能将如何使用的具体示例。例如，构建 image-editor 技能时，相关问题包括：

- "image-editor 技能应该支持什么功能？编辑、旋转，还有其他吗？"
- "能给一些这个技能如何使用的例子吗？"
- "什么样的用户输入应该触发这个技能？"

为避免让用户不堪重负，避免在单条消息中问太多问题。

### 步骤 2：规划可复用的技能内容

将具体示例转化为有效技能，通过以下方式分析每个示例：

1. 考虑如何从头执行示例
2. 识别重复执行这些工作流程时什么脚本、参考资料和资产会有帮助

### 步骤 3：分析并确定分类目录

在创建技能前，先分析技能的功能类型，确定合适的分类文件夹。参考 `.codebuddy/commands/` 的分类规则：

#### 3.1 扫描现有分类

读取 `.codebuddy/skills/` 目录，了解现有的分类文件夹。

**当前常见分类**：
- **工程/** - 工程化工具（构建、部署、发布等）
- **需求/** - 需求管理相关（排期、分析等）
- **create/** - 创建类（生成代码、文档等）
- **analyze/** - 分析类（代码分析、性能分析等）
- **automate/** - 自动化类（自动化测试、自动化操作等）
- **optimize/** - 优化类（性能优化、代码优化等）
- **get/** - 查询获取类（获取信息、查询数据等）

#### 3.2 确定分类逻辑

根据技能的**主要功能动作**确定分类：

| 功能特征 | 分类目录 | 示例 |
|---------|---------|------|
| 创建、生成、初始化 | `create/` | agent-creator, architecture-doc-generator |
| 分析、检查、诊断 | `analyze/` | code-analyzer, performance-profiler |
| 优化、改进、重构 | `optimize/` | code-optimizer, bundle-optimizer |
| 自动化操作、控制 | `automate/` | browser-control, macos-app-control |
| 查询、获取、读取 | `get/` | data-fetcher, info-getter |
| 工程化流程 | `工程/` | pre-release, ci-cd-helper |
| 需求管理 | `需求/` | requirement-scheduler, tapd-todo-planner |

**决策流程**：
1. 识别技能的**核心动作动词**（创建/分析/优化/自动化/查询/部署）
2. 匹配对应的分类目录
3. 如果技能功能跨越多个类别，选择**主要功能**的分类
4. 如果没有合适的现有分类，可以创建新分类（使用英文或中文，与现有风格一致）

#### 3.3 确定输出路径

确定分类后，skill 的完整路径为：
```
.codebuddy/skills/[分类目录]/[skill-name]/
```

**示例**：
- `agent-creator` → `.codebuddy/skills/create/agent-creator/`
- `pre-release` → `.codebuddy/skills/工程/pre-release/`
- `tapd-todo-planner` → `.codebuddy/skills/需求/tapd-todo-planner/`
- `browser-control` → `.codebuddy/skills/automate/browser-control/`

**注意**：
- 如果目标分类目录不存在，创建时会自动创建
- 分类目录名使用中文或英文，保持与现有风格一致
- 避免创建过于细分的分类，优先使用现有分类

### 步骤 4：初始化技能

从头创建新技能时，始终运行 `init_skill.py` 脚本。

**默认创建路径**：
- 技能默认创建在**用户级别**：`~/.codebuddy/skills/[skill-name]`
- 用户级别的技能对所有项目生效，可复用性更强

**如果需要创建项目级别的技能**，使用 `--path` 参数指定项目内路径：

```bash
scripts/init_skill.py <skill-name> --path .codebuddy/skills/[分类目录]
```

**示例**：
```bash
# 创建用户级别的 skill（默认，推荐）
scripts/init_skill.py browser-automation

# 创建项目级别的 skill（需要明确指定 --path）
scripts/init_skill.py browser-automation --path .codebuddy/skills/automate

# 创建项目级别的需求管理 skill
scripts/init_skill.py requirement-analyzer --path .codebuddy/skills/需求
```

脚本将：
- 在指定路径（或默认用户路径）创建技能目录
- 生成带有正确 frontmatter 和 TODO 占位符的 SKILL.md 模板
- 生成 README.md 模板
- 创建示例资源目录：`scripts/`、`references/` 和 `assets/`

**选择用户级别还是项目级别？**
- **用户级别**（推荐）：适用于通用技能，如 `macos-app-control`、`tapd-workhour-tracker` 等
- **项目级别**：仅当技能高度依赖特定项目结构或配置时使用

### 步骤 5：编辑技能

编辑技能时，记住技能是为另一个 Claude 实例使用而创建的。包含对 Claude 有益且非显而易见的信息。

#### 学习经过验证的设计模式

根据技能需求参考这些有用的指南：

- **多步骤流程**：参见 references/workflows.md 了解顺序工作流程和条件逻辑
- **特定输出格式或质量标准**：参见 references/output-patterns.md 了解模板和示例模式

#### 更新 SKILL.md

**编写指南**：
- 始终使用祈使句/不定式形式
- **必须使用中文撰写**

##### Frontmatter

编写带有 `name` 和 `description` 的 YAML frontmatter：

- `name`：技能名称（英文，kebab-case）
- `description`：**必须使用中文**。这是技能的主要触发机制，帮助 Claude 理解何时使用技能。包含技能做什么以及何时使用的具体触发器/上下文。

示例：
```yaml
---
name: docx-processor
description: 全面的文档创建、编辑和分析功能，支持修订追踪、批注、格式保留和文本提取。当 Claude 需要处理专业文档（.docx 文件）时使用：(1) 创建新文档，(2) 修改或编辑内容，(3) 处理修订追踪，(4) 添加批注，或任何其他文档任务。
---
```

##### 正文

用中文编写使用技能及其捆绑资源的指令。

#### 创建 README.md

**必须**为每个技能创建 README.md 文件，遵循上述模板格式。这是技能创建的强制要求。

##### 获取创建人信息

创建 README.md 时，**必须**通过 git 配置获取创建人信息：

```bash
git config user.name
git config user.email
```

将获取的信息填入 README.md 的创建人字段，格式为：`@用户名 (邮箱)`

示例：
```markdown
- 创建人: @yiqiuzheng (yiqiuzheng@tencent.com)
```

### 步骤 6：更新映射表

技能创建完成后，**必须**更新 `.codebuddy/skills-docs/SKILLS_COMMAND_MAP.md` 文件，在"未创建快捷指令的 Skills"表格中添加新创建的 skill 记录。

映射表格式：
```markdown
| Skill 名称 | 说明 |
|-----------|------|
| `skill-name` | skill 的简要功能说明 |
```

### 步骤 7：迭代

测试技能后，用户可能请求改进。

**迭代工作流程：**

1. 在实际任务中使用技能
2. 注意困难或低效之处
3. 确定应如何更新 SKILL.md 或捆绑资源
4. 实施更改并再次测试
