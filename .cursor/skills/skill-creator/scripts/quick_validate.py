#!/usr/bin/env python3
"""
技能快速验证脚本 - 精简版本
"""

import sys
import os
import re
import yaml
from pathlib import Path

# README.md 必需的章节
REQUIRED_README_SECTIONS = ['基本信息', '适用场景', '使用示例']

def validate_readme(skill_path):
    """验证 README.md 文件"""
    readme_path = skill_path / 'README.md'
    
    if not readme_path.exists():
        return False, "未找到 README.md - 每个技能必须包含维护文档"
    
    content = readme_path.read_text()
    
    # 检查必需的章节
    missing_sections = []
    for section in REQUIRED_README_SECTIONS:
        if section not in content:
            missing_sections.append(section)
    
    if missing_sections:
        return False, f"README.md 缺少必需章节: {', '.join(missing_sections)}"
    
    return True, "README.md 验证通过"

def validate_skill(skill_path):
    """技能基本验证"""
    skill_path = Path(skill_path)

    # 检查 SKILL.md 存在
    skill_md = skill_path / 'SKILL.md'
    if not skill_md.exists():
        return False, "未找到 SKILL.md"

    # 检查 README.md 存在并验证格式
    readme_valid, readme_message = validate_readme(skill_path)
    if not readme_valid:
        return False, readme_message

    # 读取并验证 frontmatter
    content = skill_md.read_text()
    if not content.startswith('---'):
        return False, "未找到 YAML frontmatter"

    # 提取 frontmatter
    match = re.match(r'^---\n(.*?)\n---', content, re.DOTALL)
    if not match:
        return False, "frontmatter 格式无效"

    frontmatter_text = match.group(1)

    # 解析 YAML frontmatter
    try:
        frontmatter = yaml.safe_load(frontmatter_text)
        if not isinstance(frontmatter, dict):
            return False, "frontmatter 必须是 YAML 字典"
    except yaml.YAMLError as e:
        return False, f"frontmatter 中的 YAML 无效: {e}"

    # 定义允许的属性
    ALLOWED_PROPERTIES = {'name', 'description', 'license', 'allowed-tools', 'metadata'}

    # 检查意外的属性（排除 metadata 下的嵌套键）
    unexpected_keys = set(frontmatter.keys()) - ALLOWED_PROPERTIES
    if unexpected_keys:
        return False, (
            f"SKILL.md frontmatter 中存在意外的键: {', '.join(sorted(unexpected_keys))}。"
            f"允许的属性有: {', '.join(sorted(ALLOWED_PROPERTIES))}"
        )

    # 检查必需字段
    if 'name' not in frontmatter:
        return False, "frontmatter 中缺少 'name'"
    if 'description' not in frontmatter:
        return False, "frontmatter 中缺少 'description'"

    # 提取 name 进行验证
    name = frontmatter.get('name', '')
    if not isinstance(name, str):
        return False, f"name 必须是字符串，得到 {type(name).__name__}"
    name = name.strip()
    if name:
        # 检查命名约定（连字符格式：小写字母加连字符）
        if not re.match(r'^[a-z0-9-]+$', name):
            return False, f"名称 '{name}' 应为连字符格式（仅限小写字母、数字和连字符）"
        if name.startswith('-') or name.endswith('-') or '--' in name:
            return False, f"名称 '{name}' 不能以连字符开头/结尾或包含连续连字符"
        # 检查名称长度（根据规范最多 64 个字符）
        if len(name) > 64:
            return False, f"名称过长（{len(name)} 个字符）。最大为 64 个字符。"

    # 提取并验证 description
    description = frontmatter.get('description', '')
    if not isinstance(description, str):
        return False, f"description 必须是字符串，得到 {type(description).__name__}"
    description = description.strip()
    if description:
        # 检查尖括号
        if '<' in description or '>' in description:
            return False, "description 不能包含尖括号（< 或 >）"
        # 检查描述长度（根据规范最多 1024 个字符）
        if len(description) > 1024:
            return False, f"描述过长（{len(description)} 个字符）。最大为 1024 个字符。"

    return True, "技能验证通过！"

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("用法: python quick_validate.py <skill_directory>")
        sys.exit(1)
    
    valid, message = validate_skill(sys.argv[1])
    print(message)
    sys.exit(0 if valid else 1)