def isIsomorphic(s: str, t: str) -> bool:
    if len(s) != len(t):
        return False

    t_map = {}
    s_map = {}

    for i in range(0, len(s)-1):
        s_map[s[i]] = t[i]
        t_map[t[i]] = s[i]

    # for i in range(0, len(s)-1):

    print(s_map.values() == t_map.keys())


s = "egg"
t = "add"
isIsomorphic(s, t)
