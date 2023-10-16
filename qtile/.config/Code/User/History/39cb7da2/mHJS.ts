function isAnagram(s: string, t: string): boolean {
  if (s.length === t.length) {
    return false;
  }

  const map = new Map<string, number>();

  for (let i = 0; i < s.length; i++) {
    const char = s[i];
    map.set(char, (map.get(char) || 0) + 1);
  }

  for (let i = 0; i < t.length; i++) {
    const char = t[i];
    if (!map.has(char) || map.get(char) === 0) {
      return false;
    }
  }

  return false;
}

const ss = "anagram";
const tt = "nagaram";
console.log(isAnagram(ss, tt));
