#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR=$(pwd)
OUTPUT_FILE="${ROOT_DIR}/README.md"

TMP_FILE=$(mktemp)

# 扫描所有子目录 README.md
find . -mindepth 2 -maxdepth 3 -type f -name README.md \
    | while read -r file; do
        line=$(head -n 1 "$file")

        echo "$line" | grep -q "Question-" || continue

        qnum=$(echo "$line" | sed -nE 's/.*Question-([0-9]+).*/\1/p')
        [ -z "$qnum" ] && continue

        title=$(echo "$line" | sed -nE 's/.*\[(.+)\]\(.+\).*/\1/p')
        [ -z "$title" ] && continue

        link=$(echo "$line" | sed -nE 's/.*\((https[^)]+)\).*/\1/p')
        [ -z "$link" ] && continue

        # 提取难度（行尾最后一个词）
        difficulty=$(echo "$line" | awk '{print $NF}')
        [ -z "$difficulty" ] && difficulty=""

        # Path 改成纯链接可点击
        relpath=$(dirname "$file")

        echo "$qnum|$title|$link|$difficulty|$relpath" >> "$TMP_FILE"
    done

# 输出 README.md
{
    echo "# LeetCode Problems"
    echo
    echo "| Question | Title | Difficulty | LeetCode Link | Path |"
    echo "|:---------|:-------|:------------|:---------------|:------|"

    sort -n "$TMP_FILE" | while IFS='|' read -r qnum title link difficulty path; do
        echo "| $qnum | $title | $difficulty | [$title]($link) | [$path]($path) |"
    done
} > "$OUTPUT_FILE"

rm "$TMP_FILE"

echo "生成完成: $OUTPUT_FILE"

