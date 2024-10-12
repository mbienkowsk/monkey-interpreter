let map = fn(arr, func) {
  let mapped = [];

  let acc = fn(arr, func, buf) {
    if (len(arr) == 0) {buf} else {
      let buf = push(buf, func(arr[0]));
      return acc(rest(arr), func, buf);
    }
  }

  return acc(arr, func, mapped)
}

let double = fn(x) {2 * x}

puts(map([1, 2, 3], double))
