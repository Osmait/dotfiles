def isIsomorphic(s: str, t: str) -> bool:
    if len(s) != len(t):
        return False

    t_map = {}
    s_map = {}

    for i in range(0, len(s)-1):
        s_map[s[i]] = t[i]
        t_map[t[i]] = s[i]
    print(len(t_map))


s = "foo"
t = "bar"
isIsomorphic(s, t)
