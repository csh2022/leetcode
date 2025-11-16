#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR=$(pwd)
OUTPUT_FILE="${ROOT_DIR}/README.md"

TMP_FILE=$(mktemp)

# 扫描所有子目录 README.md
find . -mindepth 2 -maxdepth 3 -type f -name README.md \
    | while read -r file; do
        line=$(head -n 1 "$file")

        echo "$line" | grep -q "Problem-" || continue

        pnum=$(echo "$line" | sed -nE 's/.*Problem-([0-9]+).*/\1/p')
        [ -z "$pnum" ] && continue

        title=$(echo "$line" | sed -nE 's/.*\[(.+)\]\(.+\).*/\1/p')
        [ -z "$title" ] && continue

        link=$(echo "$line" | sed -nE 's/.*\((https[^)]+)\).*/\1/p')
        [ -z "$link" ] && continue

        # 提取难度
        difficulty=$(echo "$line" | awk '{print $NF}')
        [ -z "$difficulty" ] && difficulty=""

        # 路径链接
        relpath=$(dirname "$file")

        # Title 和 Link 合并为 Markdown 链接
        title_link="[$title]($link)"

        echo "$pnum|$title_link|$difficulty|$relpath" >> "$TMP_FILE"
    done

# 输出 README.md
{
    echo "# LeetCode Problems"
    echo
    echo "| No. | Problem | Difficulty | Path |"
    echo "|:----|:---------|:------------|:------|"

    idx=1
    sort -n "$TMP_FILE" | while IFS='|' read -r pnum title_link difficulty path; do
        # title_link 已经是格式如: [Two Sum](https://leetcode.com/xxx)
        # 取出标题文本
        title_text=$(echo "$title_link" | sed -nE 's/\[(.+)\]\(.+\)/\1/p')

        # 生成 Problem = "pnum. title"
        problem_markdown="$pnum. [$title_text]$(echo "$title_link" | sed -nE 's/\[.+\](\(.*\))/\1/p')"

        echo "| $idx | $problem_markdown | $difficulty | [$path]($path) |"
        idx=$((idx + 1))
    done
} > "$OUTPUT_FILE"

rm "$TMP_FILE"

echo "README.md generated: $OUTPUT_FILE"

