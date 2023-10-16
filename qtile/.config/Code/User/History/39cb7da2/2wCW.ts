function isAnagram(s: string, t: string): boolean {
  if (s.length === t.length) {
    return false;
  }

  const map = new Map<string, number>();

  for (let i = 0; i < s.length; i++) {
    const char = s[i];
    map.set(char, (map.get(char) || 0) + 1);
  }
}

const s = "anagram";
const t = "nagaram";
console.log(isAnagram(s, t));
