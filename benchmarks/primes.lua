---Find primes smaller than this number:
LIMIT=100000

primes={}

---Checks to see if a number is prime or not
function isPrime(n)
	for i, prime in pairs(primes) do
		if prime*2 > n then break end
		if (n%prime == 0) then return end
	end

	table.insert(primes, n)
end

---Print all the prime numbers within range
for i=2,LIMIT do
	isPrime(i)
end

print(#primes)