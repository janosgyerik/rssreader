#!/usr/bin/env bash

urlize() {
    # not the smartest regex, but works in both GNU and BSD sed
    sed -e 's!http[s]*://[^); ]*!<a href="&">&</a>!g'
}

echo Content-type: text/html
echo "Refresh: 300;$REQUEST_URI"
echo

title="Summary ($(date))"

cat << EOF
<!doctype html>
<html>
<head>
<title>$title</title>
</head>
<body>

<h1>$title</h1>

EOF

for file in $(ls logs); do
    file=logs/$file
    printf '<h3><a href="%s">%s</a></h3>\n\n' "$file" "$file"

    echo '<pre>'
    tail -n 200 "$file" | urlize
    echo '</pre>'
done

cat << EOF

</body>
</html>
EOF
