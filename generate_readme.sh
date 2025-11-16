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
    echo "| Problem | Title | Difficulty | Path |"
    echo "|:---------|:-------|:------------|:------|"

    sort -n "$TMP_FILE" | while IFS='|' read -r pnum title_link difficulty path; do
        echo "| $pnum | $title_link | $difficulty | [$path]($path) |"
    done
} > "$OUTPUT_FILE"

rm "$TMP_FILE"

echo "生成完成: $OUTPUT_FILE"

