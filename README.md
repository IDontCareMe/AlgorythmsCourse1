![version](https://img.shields.io/badge/Go%20-v1.14-blue)

# Algorithms Course 1
Learning course from https://stepik.org/course/217 
#### Learning Go.

## The first chapter contains the fibanocci numbers algorythm and GCD
## The second chapter is about Greedy algorithm
1) По данным n отрезкам необходимо найти множество точек минимального размера, для которого каждый из отрезков содержит хотя бы одну из точек. В первой строке дано число 1 ≤ n ≤ 100 отрезков. Каждая из последующих n строк содержит по два числа 0 ≤  l ≤ r ≤ 10^9, задающих начало и конец отрезка. Выведите оптимальное число m точек и сами m точек. Если таких множеств точек несколько, выведите любое из них.
2) Первая строка содержит количество предметов 1 ≤ n ≤ 10^3 и вместимость рюкзака 0 ≤ W ≤ 2⋅10^6. Каждая из следующих nn строк задаёт стоимость 0 ≤ ci ≤2⋅10^6 и объём  0 < wi ≤2⋅10^6 предмета (n, W, ci, wi — целые числа). Выведите максимальную стоимость частей предметов (от каждого предмета можно отделить любую часть, стоимость и объём при этом пропорционально уменьшатся), помещающихся в данный рюкзак, с точностью не менее трёх знаков после запятой.
3) По данному числу 1 ≤ n ≤ 10^9 найдите максимальное число k, для которого n можно представить как сумму k различных натуральных слагаемых. Выведите в первой строке число k, во второй — k слагаемых.
### Huffman code
4) По данной непустой строке s длины не более 10^4, состоящей из строчных букв латинского алфавита, постройте оптимальный беспрефиксный код. В первой строке выведите количество различных букв k, встречающихся в строке, и размер получившейся закодированной строки. В следующих k строках запишите коды букв в формате "letter: code". В последней строке выведите закодированную строку.
5) Восстановите строку по её коду и беспрефиксному коду символов. 
В первой строке входного файла заданы два целых числа k и l через пробел — количество различных букв, встречающихся в строке, и размер получившейся закодированной строки, соответственно. В следующих k строках записаны коды букв в формате "letter: code". Ни один код не является префиксом другого. Буквы могут быть перечислены в любом порядке. В качестве букв могут встречаться лишь строчные буквы латинского алфавита; каждая из этих букв встречается в строке хотя бы один раз. Наконец, в последней строке записана закодированная строка. Исходная строка и коды всех букв непусты. Заданный код таков, что закодированная строка имеет минимальный возможный размер.
В первой строке выходного файла выведите строку s. Она должна состоять из строчных букв латинского алфавита. Гарантируется, что длина правильного ответа не превосходит 10^4 символов.
### Priority queue
6) Первая строка входа содержит число операций 1 ≤ n ≤ 10^5. Каждая из последующих `n` строк задают операцию одного из следующих двух типов:
   
`Insert x`, где 0 ≤ x ≤ 10^9 — целое число;

`ExtractMax`.

Первая операция добавляет число `x` в очередь с приоритетами, вторая — извлекает максимальное число и выводит его.