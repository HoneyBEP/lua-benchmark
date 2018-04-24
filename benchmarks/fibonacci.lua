---Find fibonacci numbers smaller than this number:
LIMIT=50

fib={}

---Calculate next fibbonacci number
function nextFibonacci()
	table.insert(fib, fib[#fib-1] + fib[#fib])
end

table.insert(fib, 1)
table.insert(fib, 1)

---Print the first LIMIT fibonacci numbers
for i=1,LIMIT do
	nextFibonacci()
end

print(fib[#fib])