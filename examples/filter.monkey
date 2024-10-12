let filter = fn(arr, cond) {
  let acc = fn(arr, func, buf) {
    if (len(arr) == 0) {buf} else {
      if (func(arr[0])) {
        let buf = push(buf, arr[0]);
      }
      return acc(rest(arr), func, buf);
    }
  }
  return acc(arr, cond, [])
}

let nums = [1, 2, 3, 4, 5, 6];
let isOdd = fn(x) {return x % 2 == 1};
puts(filter(nums, isOdd));
