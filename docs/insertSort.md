## 插入排序
- ### 基本思想
> 通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应的位置并插入。

  对于少量元素的排序，它是一个有效的算法。
  插入排序的工作方式非常类似于整理扑克牌

  ![image](http://i.niupic.com/images/2017/09/24/w7FqxU.jpg
  )

  ``` javascript
  let arr = [666, -6, 66, 66666, 6565, 89];
  for (let i = 0; i < arr.length; i++) {
  	let index = 0;
  	for (let j = 1; j < arr.length - i; j++) {
  		if (arr[j] > arr[index]) {
  			index = j;
  		}
  	}
  	let temp = arr[arr.length - i - 1];
  	arr[arr.length - i - 1] = arr[index];
  	arr[index] = temp;
  };
  console.log(arr);
  ```
