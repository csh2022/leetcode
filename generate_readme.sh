#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR=$(pwd)
OUTPUT_FILE="${ROOT_DIR}/README.md"

TMP_FILE=$(mktemp)

# 找到子目录 README.md（不再排除 template）
find . -mindepth 2 -maxdepth 3 -type f -name README.md \
    | while read -r file; do
        # 读取第一行
        line=$(head -n 1 "$file")

        # 必须包含 Question-
        echo "$line" | grep -q "Question-" || continue

        # 提取题号
        qnum=$(echo "$line" | sed -nE 's/.*Question-([0-9]+).*/\1/p')
        [ -z "$qnum" ] && continue

        # 提取标题
        title=$(echo "$line" | sed -nE 's/.*\[(.+)\]\(.+\).*/\1/p')
        [ -z "$title" ] && continue

        # 提取链接
        link=$(echo "$line" | sed -nE 's/.*\((https[^)]+)\).*/\1/p')
        [ -z "$link" ] && continue

        # 相对路径
        relpath=$(dirname "$file")

        # 写入临时文件
        echo "$qnum|$title|$link|$relpath" >> "$TMP_FILE"
    done

# 输出根 README.md
{
    echo "# LeetCode Problems"
    echo
    echo "自动生成问题总览："
    echo
    echo "| Question | Title | Link | Path |"
    echo "|---------:|--------|-------|------|"

    sort -n "$TMP_FILE" | while IFS='|' read -r qnum title link path; do
        echo "| $qnum | $title | [$title]($link) | \`$path\` |"
    done
} > "$OUTPUT_FILE"

rm "$TMP_FILE"

echo "生成完成: $OUTPUT_FILE"

