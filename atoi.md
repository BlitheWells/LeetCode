[字符串转换整数](https://leetcode-cn.com/problems/string-to-integer-atoi/)

利用状态机求解

| |' '|+/-|0-9|other|
|:--|:--|:--|:--|:--|
|start|start |signed| int_number| end|
|signed|end   |end|    int_number |end|
|int_number|end|   end    |int_number |end|
|end|end   |end|    end        |end|
