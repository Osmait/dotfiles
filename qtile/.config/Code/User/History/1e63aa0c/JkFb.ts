function isSubsequence(s: String, t: String): boolean {
  let sIndex = 0;
  let tIndex = 0;

  while (sIndex < s.length && tIndex < t.length) {
    if (s.charAt(sIndex) == t.charAt(tIndex)) {
      sIndex++;
    }
    tIndex++;
  }
  return sIndex === s.length;
}

let s = "abc";
let t = "ahbgdc";

console.log(isSubsequence(s, t));
