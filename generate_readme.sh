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

        # 提取最后两个字段
        last_word=$(echo "$line" | awk '{print $NF}')
        second_last_word=$(echo "$line" | awk '{print $(NF-1)}')

        difficulty=""
        status=""

        # 难度：只有 Easy Medium Hard
        if [[ "$second_last_word" =~ ^(Easy|Medium|Hard)$ ]]; then
            difficulty="$second_last_word"
        fi

        # 状态：只有 Done / Doing
        if [[ "$last_word" =~ ^(Done|Doing)$ ]]; then
            status="$last_word"
        fi

        # 路径链接
        relpath=$(dirname "$file")

        # Title 和 Link 合并为 Markdown 链接
        title_link="[$title]($link)"

        echo "$pnum|$title_link|$difficulty|$status|$relpath" >> "$TMP_FILE"
    done

# 输出 README.md
{
    echo "# LeetCode Problems"
    echo
    echo "| No. | Problem | Difficulty | Status | Path |"
    echo "|:----|:---------|:------------|:--------|:------|"

    idx=1
    sort -n "$TMP_FILE" | while IFS='|' read -r pnum title_link difficulty status path; do
        title_text=$(echo "$title_link" | sed -nE 's/\[(.+)\]\(.+\)/\1/p')

        problem_markdown="$pnum. [$title_text]$(echo "$title_link" | sed -nE 's/\[.+\](\(.*\))/\1/p')"

        # ★ 在这里追加固定字符串
        echo "| $idx | $problem_markdown | $difficulty | $status | [$path](${path}/problems/problem.go) |"

        idx=$((idx + 1))
    done
} > "$OUTPUT_FILE"

rm "$TMP_FILE"

echo "README.md generated: $OUTPUT_FILE"

