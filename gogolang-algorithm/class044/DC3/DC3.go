package main

// sa 数组
func sa(nums []int, max int) []int {
	n := len(nums)
	arr := make([]int, n+3)
	for i := 0; i < n; i++ {
		arr[i] = nums[i]
	}
	return skew(arr, n, max)
}

func skew(nums []int, n, K int) []int {
	n0 := (n + 2) / 3
	n1 := (n + 1) / 3
	n2 := n / 3
	n02 := n0 + n2
	s12 := make([]int, n02+3)
	sa12 := make([]int, n02+3)
	for i, j := 0, 0; i < n+(n0-n1); i++ {
		if 0 != i%3 {
			s12[j] = i
			j++
		}
	}
	radixPass(nums, s12, sa12, 2, n02, K)
	radixPass(nums, sa12, s12, 1, n02, K)
	radixPass(nums, s12, sa12, 0, n02, K)
	name := 0
	c0 := -1
	c1 := -1
	c2 := -1
	for i := 0; i < n02; i++ {
		if c0 != nums[sa12[i]] || c1 != nums[sa12[i]+1] || c2 != nums[sa12[i]+2] {
			name++
			c0 = nums[sa12[i]]
			c1 = nums[sa12[i]+1]
			c2 = nums[sa12[i]+2]
		}
		if 1 == sa12[i]%3 {
			s12[sa12[i]/3] = name
		} else {
			s12[sa12[i]/3+n0] = name
		}
	}
	if name < n02 {
		sa12 = skew(s12, n02, name)
		for i := 0; i < n02; i++ {
			s12[sa12[i]] = i + 1
		}
	} else {
		for i := 0; i < n02; i++ {
			sa12[s12[i]-1] = i
		}
	}
	s0 := make([]int, n0)
	sa0 := make([]int, n0)
	for i, j := 0, 0; i < n02; i++ {
		if sa12[i] < n0 {
			s0[j] = 3 * sa12[i]
			j++
		}
	}
	radixPass(nums, s0, sa0, 0, n0, K)
	sa := make([]int, n)
	for p, t, k := 0, n0-n1, 0; k < n; k++ {
		i := 0
		if sa12[t] < n0 {
			i = sa12[t]*3 + 1
		} else {
			i = (sa12[t]-n0)*3 + 2
		}
		j := sa0[p]

		var tempInsert bool
		if sa12[t] < n0 {
			tempInsert = leq(nums[i], s12[sa12[t]+n0], nums[j], s12[j/3])
		} else {
			tempInsert = leq2(nums[i], nums[i+1], s12[sa12[t]-n0+1], nums[j], nums[j+1], s12[j/3+n0])
		}
		if tempInsert {
			sa[k] = i
			t++
			if t == n02 {
				for k++; p < n0; p++ {
					sa[k] = sa0[p]
					k++
				}
			}
		} else {
			sa[k] = j
			p++
			if p == n0 {
				for k++; t < n02; t++ {
					if sa12[t] < n0 {
						sa[k] = sa12[t]*3 + 1
					} else {
						sa[k] = (sa12[t]-n0)*3 + 2
					}
					k++
				}
			}
		}
	}
	return sa
}

func radixPass(nums, input, output []int, offset, n, k int) {
	cnt := make([]int, k+1)
	for i := 0; i < n; i++ {
		cnt[nums[input[i]+offset]]++
	}
	for i, sum := 0, 0; i < len(cnt); i++ {
		t := cnt[i]
		cnt[i] = sum
		sum += t
	}
	for i := 0; i < n; i++ {
		output[cnt[nums[input[i]+offset]]] = input[i]
		cnt[nums[input[i]+offset]]++
	}
}

// rand 数组
// 改写了一下，将生成的 sa数组作为参数传入到函数中
func rank(sa []int) []int {
	n := len(sa)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[sa[i]] = i
	}
	return ans
}

// height 数组
// 改写了一下，将生成的 sa,rank数组作为参数传入到函数中
func height(s []int, sa, rank []int) []int {
	n := len(s)
	ans := make([]int, n)
	for i, k := 0, 0; i < n; i++ {
		if rank[i] != 0 {
			if k > 0 {
				k--
			}
			j := sa[rank[i]-1]
			for i+k < n && j+k < n && s[i+k] == s[j+k] {
				k++
			}
			ans[rank[i]] = k
		}
	}
	return ans
}

func main() {

}

// lexicographic order
func leq(a1, a2, b1, b2 int) bool {
	return a1 < b1 || (a1 == b1 && a2 <= b2) // for pairs
}

func leq2(a1, a2, a3, b1, b2, b3 int) bool {
	return a1 < b1 || (a1 == b1 && leq(a2, a3, b2, b3)) // and triples
}

//// stably sort a[0..n-1] to b[0..n-1] with keys in 0..K from r
//func radixPass(a, b, r []int, n, K int) { // count occurrences
//	c := make([]int, K+1) // counter array
//	for i := 0; i <= K; i++ {
//		c[i] = 0 // reset counters
//	}
//	for i := 0; i < n; i++ {
//		c[r[a[i]]]++ // count occurrences
//	}
//	for i, sum := 0, 0; i <= K; i++ { // exclusive prefix sums
//		t := c[i]
//		c[i] = sum
//		sum += t
//	}
//	for i := 0; i < n; i++ {
//		b[c[r[a[i]]]] = a[i] // sort
//		c[r[a[i]]]++
//	}
//}

/*

inline bool leq(int a1, int a2, int b1, int b2) // lexicographic order
{return(a1 < b1 || a1 == b1 && a2 <= b2); } // for pairs
inline bool leq(int a1, int a2, int a3, int b1, int b2, int b3)
{return(a1 < b1 || a1 == b1 && leq(a2,a3, b2,b3)); } // and triples


// stably sort a[0..n-1] to b[0..n-1] with keys in 0..K from r
static void radixPass(int* a, int* b, int* r, int n, int K)
{// count occurrences
int* c = new int[K + 1]; // counter array
for (int i = 0; i <= K; i++) c[i] = 0; // reset counters
for (int i = 0; i < n; i++) c[r[a[i]]]++; // count occurrences
for (int i = 0, sum = 0; i <= K; i++) // exclusive prefix sums
{int t = c[i]; c[i] = sum; sum += t; }

for (int i = 0; i < n; i++) b[c[r[a[i]]]++] = a[i]; // sort
delete [] c;
}

// find the suffix array SA of s[0..n-1] in {1..K}ˆn
// require s[n]=s[n+1]=s[n+2]=0, n>=2
void suffixArray(int* s, int* SA, int n, int K) {
int n0=(n+2)/3, n1=(n+1)/3, n2=n/3, n02=n0+n2;
int* s12 = new int[n02 + 3]; s12[n02]= s12[n02+1]= s12[n02+2]=0;
int* SA12 = new int[n02 + 3]; SA12[n02]=SA12[n02+1]=SA12[n02+2]=0;
int* s0 = new int[n0];
int* SA0 = new int[n0];
// generate positions of mod 1 and mod 2 suffixes
// the "+(n0-n1)" adds a dummy mod 1 suffix if n%3 == 1
for (int i=0, j=0; i < n+(n0-n1); i++) if (i%3 != 0) s12[j++] = i;
// lsb radix sort the mod 1 and mod 2 triples

radixPass(s12 , SA12, s+2, n02, K);
radixPass(SA12, s12 , s+1, n02, K);
radixPass(s12 , SA12, s , n02, K);

// find lexicographic names of triples
int name = 0, c0 = -1, c1 = -1, c2 = -1;
for (int i = 0; i < n02; i++) {
if (s[SA12[i]] != c0 || s[SA12[i]+1] != c1 || s[SA12[i]+2] != c2)
{name++; c0 = s[SA12[i]]; c1 = s[SA12[i]+1]; c2 = s[SA12[i]+2]; }
if (SA12[i] % 3 == 1) { s12[SA12[i]/3] = name; } // left half
else {s12[SA12[i]/3 + n0] = name; } // right half
}
// recurse if names are not yet unique
if (name < n02) {
suffixArray(s12, SA12, n02, name);
// store unique names in s12 using the suffix array
for (int i = 0; i < n02; i++) s12[SA12[i]] = i + 1;
} else // generate the suffix array of s12 directly
for (int i = 0; i < n02; i++) SA12[s12[i] - 1] = i;
// stably sort the mod 0 suffixes from SA12 by their first character
for (int i=0, j=0; i < n02; i++) if (SA12[i] < n0) s0[j++] = 3*SA12[i];
radixPass(s0, SA0, s, n0, K);
// merge sorted SA0 suffixes and sorted SA12 suffixes
for (int p=0, t=n0-n1, k=0; k < n; k++) {
#define GetI() (SA12[t] < n0 ? SA12[t]*3+1: (SA12[t] - n0) * 3 + 2)
int i = GetI(); // pos of current offset 12 suffix
int j = SA0[p]; // pos of current offset 0 suffix
if (SA12[t] < n0 ? // different compares for mod 1 and mod 2 suffixes
leq(s[i], s12[SA12[t] + n0], s[j], s12[j/3]) :

leq(s[i],s[i+1],s12[SA12[t]-n0+1], s[j],s[j+1],s12[j/3+n0]))
{// suffix from SA12 is smaller
SA[k] = i; t++;
if (t == n02) // done --- only SA0 suffixes left
for (k++; p < n0; p++, k++) SA[k] = SA0[p];
} else {// suffix from SA0 is smaller
SA[k] = j; p++;
if (p == n0) // done --- only SA12 suffixes left
for (k++; t < n02; t++, k++) SA[k] = GetI();
} }
delete [] s12; delete [] SA12; delete [] SA0; delete [] s0;
}

*/
