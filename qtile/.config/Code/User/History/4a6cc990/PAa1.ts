function isIsomorpic(s: String, t: String): boolean {
  if (s.length !== t.length) return false;

    const sMap = new Map<String, String>();
    const tMap = new Map<String, String>();

    for (let i = 0; i < s.length; i++) {
        let c1 = s.charAt(i);
        let c2 = t.charAt(i);
        if((tMap.has(c2) && tMap[c1] !== c2))
    }
}
