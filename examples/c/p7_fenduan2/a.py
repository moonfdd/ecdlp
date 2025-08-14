import math

# fillPrime function fills primes from 2 to sqrt of high in chprime list
def fillPrimes(chprime, high):

	ck = [True]*(high+1)

	l = int(math.sqrt(high))
	for i in range(2, l+1):
		if ck[i]:
			for j in range(i*i, l+1, i):
				ck[j] = False
				
	for k in range(2, l+1):
		if ck[k]:
			chprime.append(k)

# in segmented sieve we check for prime from range [low, high]
def segmentedSieve(low, high):
	
	chprime = list()
	fillPrimes(chprime, high)
# chprimes has primes in range [2,sqrt(n)]
# we take primes from 2 to sqrt[n] because the multiples of all primes below high are marked by these 
# primes in range 2 to sqrt[n] for eg: for number 7 its multiples i.e 14 is marked by 2, 21 is marked by 3,
# 28 is marked by 4, 35 is marked by 5, 42 is marked 6, so 49 will be first marked by 7 so all number before 49 
# are marked by primes in range [2,sqrt(49)] 
	prime = [True] * (high-low + 1)
# here prime[0] indicates whether low is prime or not similarly prime[1] indicates whether low+1 is prime or not
	for i in chprime:
		lower = (low//i)
# here lower means the first multiple of prime which is in range [low,high]
# for eg: 3's first multiple in range [28,40] is 30		 
		if lower <= 1:
			lower = i+i
		elif (low % i) != 0:
			lower = (lower * i) + i
		else:
			lower = lower*i
		for j in range(lower, high+1, i):
			prime[j-low] = False
			
			
	for k in range(low, high + 1):
			if prime[k-low]:
				print(k, end=" ")


#DRIVER CODE
# low should be greater than or equal to 2
low = 2

# low here is the lower limit
high = 100

# high here is the upper limit
# in segmented sieve we calculate primes in range [low,high] 
# here we initially we find primes in range [2,sqrt(high)] to mark all their multiples as not prime
# then we mark all their(primes) multiples in range [low,high] as false
# this is a modified sieve of eratosthenes , in standard sieve of eratosthenes we find prime from 1 to n(given number)
# in segmented sieve we only find primes in a given interval 
print("Primes in Range %d to %d are" )
segmentedSieve(low, high)
