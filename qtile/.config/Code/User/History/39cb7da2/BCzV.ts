function isAnagram(s: string, t: string): boolean {
  const map1 = new Map<string, number>();
  const map2 = new Map<string, number>();
  for (let i = 0; i < s.length; i++) {
    const char = s[i];
    map1.set(char, (map1.get(char) || 0) + 1);
    map2.set(char, (map2.get(char) || 0) + 1);
  }
  return map1 === map2;
}

const s = "anagram";
const t = "nagaram";
console.log(isAnagram(s, t));
