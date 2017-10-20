# include<cstdio>
# include<queue>
using namespace std;

queue<int> q;
int main() {
	int n;
	scanf("%d", &n);
	for (int i=0; i < n; i++) {
		q.push(i + 1);
	}
	while (!q.empty()) {
		printf("%d\n", q.front());
		q.pop();
		q.push(q.front());
		q.pop();
	}
}