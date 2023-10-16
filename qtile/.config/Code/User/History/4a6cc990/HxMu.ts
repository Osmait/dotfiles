function isIsomorpic(s: String, t: String): boolean {
  if (s.length !== t.length) return false;

  const sMap = new Map<String, String>();
  const tMap = new Map<String, String>();

  for (let i = 0; i < s.length; i++) {
    let c1 = s.charAt(i);
    let c2 = t.charAt(i);
    if (
      (tMap.has(c2) && tMap.get(c2) !== c1) ||
      (sMap.has(c1) && sMap.get(c1) !== c2)
    )
      return false;
    sMap.set(c1, c2);
    tMap.set(c2, c1);
  }
  return true;
}

console.log(isIsomorpic("egg", "add"));
