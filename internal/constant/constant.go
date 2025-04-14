package constant

const (
	Version    = "v1.21.6"                                     // 版本号
	RepoAddr   = "https://github.com/AmbitiousJun/live-server" // 仓库地址
	HeadersSeg = "[[[:]]]"                                     // 分割请求头的分隔符
)

const (
	Dir_DataRoot = "data" // 数据存放根目录
	// Dir_DataRoot = "/Users/ambitious/Desktop/code/go/live-server/data" // 数据存放根目录, 开发环境
)

const (

	// HelpDocHtmlTemplate 帮助文档的 html 模板
	HelpDocHtmlTemplate = `PCFET0NUWVBFIGh0bWw+CjxodG1sIGxhbmc9ImVuIj4KICA8aGVhZD4KICAgIDxtZXRhIGNoYXJzZXQ9IlVURi04IiAvPgogICAgPG1ldGEgbmFtZT0idmlld3BvcnQiIGNvbnRlbnQ9IndpZHRoPWRldmljZS13aWR0aCwgaW5pdGlhbC1zY2FsZT0xLjAiIC8+CiAgICA8bGluayByZWw9Imljb24iIHR5cGU9ImltYWdlL3gtaWNvbiIgaHJlZj0iZGF0YTppbWFnZS92bmQubWljcm9zb2Z0Lmljb247YmFzZTY0LEFBQUJBQUVBSUNBQUFBRUFJQUNvRUFBQUZnQUFBQ2dBQUFBZ0FBQUFRQUFBQUFFQUlBQUFBQUFBQUJBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFEOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8vdjkvZi8vLy8vLy8vLy8vLy8vLy8vNi9Qei8rL3o4Ly8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLzcvZnovKy8zOS8vLy8vLy8vLy8vLy8vLy8vL3Y5L2YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvNy9mMy8vLy8vLytQazVQK1FsSlgveE1YRS8vLy8vLy8vLy8vLzdlLzIvOXZkNVAvTnp0VC94TVhLLzhIQ3h2L0J3OGYveGNmTS84ek8wLy9aMitILzdPNzEvLy8vLy8vLy8vLy92c0MvLzVPWW1mL2w1dWIvLy8vLy8vdjkvZi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8vbjcrLy8vLy8vL21acWEvMHAxay85V2Y2TC9jWEJ2LzQrUGZ2K0VnbDcvaW9oVy81Q09VditVa2xIL21KWlMvNWlWVXYrVmsxTC9rSTFTLzRtSFZQK0ZoRjMva1pGNy8yMXRhZjlWZjZQL1NuV1QvNStob1AvLy8vLy8rZnY3Ly96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8rLzM5Ly8vLy8vLzA5ZmIvV21kdi8zNkZZUCswc0Z2LzBzNXAvK1BmZGYvbzVYci82dWQ4LytybmZmL3E1MzMvNnVkOS8rcm5mZi9xNTN6LzZlVjYvK1BmZGYvVTBHdi90TEJjLzRxUlpmOWFaV24vN3UveS8vLy8vLy83L2YzLy9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi82L1B6Ly8vLy8vNTZmbmYra25sUC8zdGx4LzdtNGFmK2xvMkQvbDVOWS80MkpVditHZ1U3L2czMU0vNEY4Uy8rQWVrci9nWHhNLzRlRFQvK1FqVlgvbTVwYi83Q3VaUC9iMkhiLzVONTEvNnVsVmYrUGtJai8vLy8vLy92OS9QLzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8rL3o4Ly8vLy8vL1MxTnYvaUlaUC84L01idjlQVEQ3L09DNHUvek10TVA4eU1qZi9NRFUrL3pJNlEvOHlQVWIvTVR4Ry96QThSdjh6UFViL01UYy8vekl4T1A4MkxqRC9PVEF6LzJSaVJmL095Mi8vNU9GNy81R09UUCsrdjhULy8vLy8vL3I4L1AvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi81Ky92Ly8vLy8vNXFibC8vQ3YySC9pSVpWL3pFdk9mOUJmWTcvVDY3SS8xaS8ydjlUd2VEL1RyL2gvMXZPN1A5YXplei9Uc0xsLzFYSTZQOWF5T1QvU3JEUS8weXB3LzlIYW5ML055ODAvNnFwWVAvZTIzbi94Y0ZrLzQ2TmZ2Ly8vLy8vK3Z6OC8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy92OS9QLy8vLy8vaTR0My85UFJhLzl1WmtUL09FNVovMXJiL3Y5UnpmYi9XOVg1LzEvWit2OVd6L1QvVWNudy8xN1crUDljMWZiL1VNanYvMW5TOXY5ZTEvbi9WTmYvLzAyd3kvODJMelAvbDVWWi85M1pkdi9WMFcvL2lZZGkvKy94K1AvLy8vLy8vUDMrLy96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly92Ly8vL1gzL2YrSmlHbi8yZGR3LzE5Vk8vODZYV3ovWDl2OC8xbk43djlOdytyL1djL3gvMTdWOWY5VnpPLy9VTWZzLzEzVDlQOWIwdlAvVDhYci8xZk43LzllMXZqL1M3UFQvek0wTy8rTGlGUC8zTmwyLzluV2MvK01pMXIvNE9McC8vLy8vLy83L2YzLy9QNysvL3orL3YvOC92Ny8vUDcrLy92OS9mLy8vLy8vN083MS80aUhZZi9hMlhML1dVNDQvemRnY3Y5VTAvbi9YdFB6LzFuUTh2OU54T3YvV05EeS8xN1Y5djlVelBEL1VNZnQvMTNVOWY5YzAvVC9UOFhyLzFuUzl2OVp4K1QvTXpoQi80Ti9Udi9iMkhYLzJ0WjAvNCtOVi8vVzJOLy8vLy8vLy9yOC9QLzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9mLy8vLy8vLy8vMDlQbi9qWXhrLzlyWWN2OVZTamIvUDJwNS8xVFUrdjlTeCt6L1h0WDIvMW5ROC85TXhPdi9XTS95LzE3Vzl2OVZ6UEQvVU1mdC8xM1U5ZjliMHZQL1VNbncvMVBENC84MlBVWC9mM3BNLytMZmVQL2gzWGIvbHBSWi85emM0Zi8vLy8vLy9QNy8vL3orL3YvOC92Ny8vUDcrLy96Ky92Ly8vLy8vazVxZC8zcUZrditGZ2xqLzNOcDAvMVZLTnY4OVpuYi9ZTi8vLzFISDdQOVN5ZTcvWHRYMi8xblI4LzlOeE92L1Y4L3kvMTdXOXY5VnpQRC9VTWZ0LzF6VDlQOWQxdmovVEwzZS96STVRLytFZjAvL2hJSlIvMlZqUmYrUmpsVC9kbjZJLzVTYm52LzUrZm4vL2YvLy8veisvdi83L2YzLy8vLy8vK1RpNFA5TVpuWC9aYUxULzNsOVdmL2YyM1QvV0U0NS96ZGNiUDlVMGZmL1h0VHovMUxKN3Y5U3llNy9YdFgyLzFyUjgvOU54T3YvVjg3eC8xN1c5djlVelBEL1VNYnMvMTdYK2Y5V3c5Ly9NemRBLzRTQVQvOXhiMHIvVUUwLy80YUVUZjlqbThiL1RHdUEvOW5YMWYvLy8vLy8rL3o5Ly96Ky92LzcvZjcvLy8vLy83ckF3djlwZ3BqL2VYeFkvK0hlZGY5ZFZEdi9PMWxsLzFUVCtmOVR5TzMvWHRUMS8xTEk3ZjlReHV2L1hkUHovMXJROGY5Tncrbi9WODN3LzEvVjlmOVZ5KzcvVk5INS8xVysydjgwTVRqL2lZWlQvN2kxWlAra29Wci9tSmRhLzJGMWlmKzl3OFgvLy8vLy8vdjkvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vLy8vLytubjZ2K0RnVnYvNHQ5MS8yaGpSZjg1UUVqL1ZiN1gvMVBPOC85VDBQWC9YOTMrLzFUUytQOVUwdnIvWU43Ly8xN2MvLzlPeS9QL1Z0TDIvMXpVOVA5U3lPdi9Rb1NWL3p3ek9QK0loMUgvVlZOQi96YzBPdjl4YjBYLzJ0alovLy8vLy8vNy9mMy8vUDcrLy96Ky92LzgvdjcvL1A3Ky8vejkvdi8rLy83Lzd2SDQvNHFKWS8vaTNuWC9rWkJXL3pFc05mODNPa0gvT0U5YS96WlVZLzg3WFd2L1FtaDEvek5iYVA4d1dHUC9ObHRsLzBCZ2JQODRVMkQvTzA5Wi96aEFSLzg3TlR2L05qTTMvN0t2WWYrUGpGWC9lSFZOLzRtSFZQL1gydUgvLy8vLy8vdjgvZi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy83Ly8vLzIrUDcvZlh4ZC85ZlRhLy9ZMUhYL21KWlkvM052U2Y5bFhUNy9YRk02LzFwUlBQOCtNUjcvZzNkdy82U2FsZitTaG9ML1BUQWcvMXBTUGY5YVVqei9aRjlELzI1dFNmK2dubHIvMTlSei8rRGRkLy9oM1cvL2NuQkovK0xrNi8vLy8vLy8rLzM5Ly96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Lyt2ejgvLy8vLy8rNHViai9mSHRTLzlITmNQL2QybmYvNE4xNC8rWGlldi9aMTNELzFOQnEvMnQ1WlA5N21hLy9oSldlLzRLZXNmOWZjbTcveDhSbi85M2JjZi9qNEhmLzROMTQvOXJXZGYvWjFYVC9vcUJkLzNkMlgvK3dzYkQvLy8vLy8vdjkvZi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvK2Z2Ni8vLy8vLytWbHBiL3M3QmEvOXZZZWYvSnhtbi9qb3hWLzVHUmN2OTRjMDcvVUhLSi8zZXUxdjlKWUcvL2NhUEYvMXFHcVA5clowVC9sWlJ2LzQrTlZ2L0l4V2ovMU5GMC85ZlViLzk5ZkZ6LzZ1ejQvLy8vLy8vNy9mMy8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvNSsvdi8vLy8vLzZ5dHJQK3FwMWovMjlkei80R0FZUC9NenRmLytQci8vLy8vLy8rb3Fxci9LU2trLzBBM00vOG5KU0wvbHBpWS8vcjgvLy8xOS8zLzBOTGIvNEtCWlAvUHkyei8wODl0LzRlSGEvLzYvUC8vKy96Ny8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8vcjgvUC8vLy8vL3hjZk0vNk9oVy8rcnFGdi91YnZDLy8vLy8vLy8vLy8vdnIrOS83cTd2ditPalhuL3BhTlIvNG1JY3YvRnhzci91THE0Ly8vLy8vLy8vLy8veE1YTi81MmJXZi9SelduL2taQ0EvLy8vLy8vNy9Qei8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8rdno4Ly8vLy8vL2k1UEgvaFlSWS81MmJXLy9aMitMLy8vLy8vLy8vLy84K1Bqei9TVWxMLzZpb252KzBzRkgvcjYrZi8wcEtUZjgwTkRMLytQcjYvLy8vLy8vajVlei9uSnBrLzdDdFYvK1FrWmovLy8vLy8vbjcrLy84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8vdjkvZi8vLy8vLzZ1encvNFNFYi85M2RVai9sNVJSLzc2L3cvLy8vLy8vLy8vLy85cmIyZi9hM09IL2taRisvOGZEWHYrUWozWC8xZGJkLzlIU3ovLysvLy8vLy8vLy84TER6UCtucEZ2L2pZdFMvMjl0U2YrMXRyUC8vLy8vLy9yOC9QLzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8rdnY3Ly8vLy8vK1ptcFgvdDdOWC80Si9UdjlwWno3L2w1WnIvOVhXMy8vNSsvLy8vZi8vLzZlb3JQK1psMVgvNU9GNy82U2hWLythbTVyLytmdi8vL2Y1Ly8vVTF1SC9tSmR3LzVtV1UvOU1TamIvMHM1ci80U0NXZi9xN1BQLy8vLy8vL3o5L3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi82Ky92Ly8vLy8vNXlkbHYrcXBsSC90YkprLzQyS1RmOWRXelAvZTNsRi81V1VjUCtVazJuL3JhcGIvK0hkZC8vWDAzVC80ZDU0LzdTeFhmK1RrV1QvbDVady80ZUZUUDk2ZHovL2JHby8vNjJxWC8rK3VtTC9oNFZjLyt2dDlQLy8vLy8vL1AzKy8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy92OS9mLy8vLy8vNnV6dy80V0ViditZbFZmL2k0cHMvN3U4eFArYW1vei9mMzFULzVLUFR2K21vMWYvdGJKbC83dTRaLys1dG1mL3JxdGYvNWFUVHYrRmcxSC9pb3B4LzdxN3dmK2JuSkwvbFpKVy80U0NWUCswdGJiLy8vLy8vL3I4L1AvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9mNysvL3Y5L2YvLy8vLy8rdnovLzlYWDMvLzA5dnYvLy8vLy8vLy8vLy9xN1BULzBkTGEvNlducS84b0ppYi9OVEl4L3lRaElQOU5UVkQveTgzVS8rRGk2di8rLy8vLy8vLy8vLy8vLy8vVzJPRC80dVRyLy8vLy8vLzcvZjMvL2Y3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL2Y3Ky8vdjkvZi85Ly83Ly8vLy8vLzcvLy8vNSsvdi8rdnY3Ly83Ly92Ly8vLy8vbTV1YS80YUhodisvdjcvL3JLMnQvMnRyYXYvLy8vLy8vLy8vLy9yNysvLzUrL3YvKy8zOC8vLy8vLy8vLy8vLyt2ejgvLzMrL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzcvZjMvL1A3Ky8veisvdi84L3Y3Ly8vLy8vNVNWbFArTWpJei8vLy8vLy8vLy8vLy8vLy8vdzhURS8yOXZidi8vLy8vLy8vLy8vL3Y5L2YvOC92Ny8rLzM5Ly92OS9mLzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vZjcrLy8zKy92LzgvdjcvL1A3Ky8vMysvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDMrLy8vLy8vK2ZuNS8vbloyZC8vLy8vLy82L1B6LytmdjcvL2I0K1AvLy8vLy91YnE1LzRpSWlQLy8vLy8vL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL2Y3Ky8veisvdi85L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8vcjgvUC8vLy8vL3djTEMvNWlZbVAvLy8vLy8rZnY3Ly96Ky92LzkvdjcvL2Y3Ly8vajUrdi8vLy8vL29LR2gvOGJIeC8vLy8vLy8rdno4Ly96Ky92LzgvdjcvL1A3Ky8vMysvdi85L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy8zKy92LzgvdjcvL1A3Ky8vMysvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy8zKy8vLzcvUDMvLy8vLy8vdjkvZi84L3Y3Ly9mNysvLzMrL3YvOS92Ny8vZjcrLy92OC9QLy8vLy8vL3YvLy8veisvdi84L3Y3Ly9QNysvLzMrL3YvOC92Ny8vUDcrLy8zKy92LzkvdjcvL1A3Ky8vMysvdi85L3Y3L0FBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQT0iPgogICAgPHRpdGxlPmxpdmUtc2VydmVyIOW4ruWKqeaWh+ahozwvdGl0bGU+CiAgICA8c3R5bGU+CiAgICAgIC5jb250YWluZXIgewogICAgICAgIG92ZXJmbG93LXdyYXA6IGJyZWFrLXdvcmQ7CiAgICAgICAgd29yZC1icmVhazogYnJlYWstYWxsOwogICAgICB9CiAgICAgIC5jb250YWluZXIgYTpsaW5rLAogICAgICAuY29udGFpbmVyIGE6dmlzaXRlZCB7CiAgICAgICAgY29sb3I6ICMxYTczZTg7CiAgICAgICAgdGV4dC1kZWNvcmF0aW9uOiBub25lOwogICAgICB9CiAgICA8L3N0eWxlPgogIDwvaGVhZD4KICA8Ym9keT4KICAgIDxkaXYgY2xhc3M9ImNvbnRhaW5lciI+JHtkb2NDb250ZW50fTwvZGl2PgoKICAgIDxzY3JpcHQ+CiAgICAgIHdpbmRvdy5vbmxvYWQgPSAoKSA9PiB7CiAgICAgICAgY29uc3QgY29udGFpbmVyID0gZG9jdW1lbnQucXVlcnlTZWxlY3RvcigiLmNvbnRhaW5lciIpOwogICAgICAgIGNvbnRhaW5lci5pbm5lckhUTUwgPSBjb250YWluZXIuaW5uZXJIVE1MLnJlcGxhY2UoCiAgICAgICAgICAvXCRce2NsaWVudE9yaWdpblx9L2csCiAgICAgICAgICBsb2NhdGlvbi5vcmlnaW4KICAgICAgICApOwogICAgICB9OwogICAgPC9zY3JpcHQ+CiAgPC9ib2R5Pgo8L2h0bWw+Cg==`

	// FengAuthHtml 凤凰秀授权地址
	FengAuthHtml = `PCFET0NUWVBFIGh0bWw+CjxodG1sPgogIDxoZWFkPgogICAgPG1ldGEgY2hhcnNldD0iVVRGLTgiIC8+CiAgICA8bWV0YSBuYW1lPSJ2aWV3cG9ydCIgY29udGVudD0id2lkdGg9ZGV2aWNlLXdpZHRoLCBpbml0aWFsLXNjYWxlPTEuMCIgLz4KICAgIDx0aXRsZT7lh6Tlh7Dnp4DnmbvlvZXmjojmnYM8L3RpdGxlPgogICAgPCEtLSDlvJXlhaUgQm9vdHN0cmFwIENTUyAtLT4KICAgIDxsaW5rCiAgICAgIGhyZWY9Imh0dHBzOi8vZmFzdGx5LmpzZGVsaXZyLm5ldC9ucG0vYm9vdHN0cmFwQDUuMy4wL2Rpc3QvY3NzL2Jvb3RzdHJhcC5taW4uY3NzIgogICAgICByZWw9InN0eWxlc2hlZXQiCiAgICAvPgogICAgPHN0eWxlPgogICAgICAucmVzdWx0LXdyYXAgewogICAgICAgIG1hcmdpbi10b3A6IDQwcHg7CiAgICAgICAgcGFkZGluZzogMjBweDsKICAgICAgfQogICAgICAucmVzdWx0LXdyYXAgLm1lc3NhZ2UgewogICAgICAgIHBhZGRpbmc6IDEwcHg7CiAgICAgIH0KICAgIDwvc3R5bGU+CiAgPC9oZWFkPgogIDxib2R5PgogICAgPGRpdiBjbGFzcz0iY29udGFpbmVyIG10LTQiPgogICAgICA8IS0tIOaJi+acuuWPt+i+k+WFpeahhiAtLT4KICAgICAgPGRpdiBjbGFzcz0ibWItMyByb3ciPgogICAgICAgIDxsYWJlbCBmb3I9InBob25lIiBjbGFzcz0iY29sLXNtLTIgY29sLWZvcm0tbGFiZWwiPuaJi+acuuWPtzwvbGFiZWw+CiAgICAgICAgPGRpdiBjbGFzcz0iY29sLXNtLTEwIj4KICAgICAgICAgIDxpbnB1dAogICAgICAgICAgICB0eXBlPSJ0ZXh0IgogICAgICAgICAgICBjbGFzcz0iZm9ybS1jb250cm9sIgogICAgICAgICAgICBpZD0icGhvbmUiCiAgICAgICAgICAgIHBsYWNlaG9sZGVyPSLor7fovpPlhaXmiYvmnLrlj7ciCiAgICAgICAgICAvPgogICAgICAgIDwvZGl2PgogICAgICA8L2Rpdj4KCiAgICAgIDwhLS0g5a+G56CB6L6T5YWl5qGGIC0tPgogICAgICA8ZGl2IGNsYXNzPSJtYi0zIHJvdyI+CiAgICAgICAgPGxhYmVsIGZvcj0icGFzc3dvcmQiIGNsYXNzPSJjb2wtc20tMiBjb2wtZm9ybS1sYWJlbCI+5a+G56CBPC9sYWJlbD4KICAgICAgICA8ZGl2IGNsYXNzPSJjb2wtc20tMTAiPgogICAgICAgICAgPGlucHV0CiAgICAgICAgICAgIHR5cGU9InBhc3N3b3JkIgogICAgICAgICAgICBjbGFzcz0iZm9ybS1jb250cm9sIgogICAgICAgICAgICBpZD0icGFzc3dvcmQiCiAgICAgICAgICAgIHBsYWNlaG9sZGVyPSLor7fovpPlhaXlr4bnoIEiCiAgICAgICAgICAvPgogICAgICAgIDwvZGl2PgogICAgICA8L2Rpdj4KCiAgICAgIDwhLS0g5Zyw5Yy656CB6L6T5YWl5qGGIC0tPgogICAgICA8ZGl2IGNsYXNzPSJtYi0zIHJvdyI+CiAgICAgICAgPGxhYmVsIGZvcj0iY29kZSIgY2xhc3M9ImNvbC1zbS0yIGNvbC1mb3JtLWxhYmVsIj7lnLDljLrnoIE8L2xhYmVsPgogICAgICAgIDxkaXYgY2xhc3M9ImNvbC1zbS0xMCI+CiAgICAgICAgICA8aW5wdXQKICAgICAgICAgICAgdHlwZT0idGV4dCIKICAgICAgICAgICAgY2xhc3M9ImZvcm0tY29udHJvbCIKICAgICAgICAgICAgaWQ9ImNvZGUiCiAgICAgICAgICAgIHBsYWNlaG9sZGVyPSLor7fovpPlhaXlnLDljLrnoIEiCiAgICAgICAgICAgIHZhbHVlPSI4NiIKICAgICAgICAgIC8+CiAgICAgICAgPC9kaXY+CiAgICAgIDwvZGl2PgoKICAgICAgPCEtLSDlh6Tlh7Dnp4Dku6TniYwgLS0+CiAgICAgIDxkaXYgY2xhc3M9Im1iLTMgcm93Ij4KICAgICAgICA8bGFiZWwgZm9yPSJmZW5nX3Nob3dfdG9rZW4iIGNsYXNzPSJjb2wtc20tMiBjb2wtZm9ybS1sYWJlbCIKICAgICAgICAgID7lh6Tlh7Dnp4Dku6TniYw8L2xhYmVsCiAgICAgICAgPgogICAgICAgIDxkaXYgY2xhc3M9ImNvbC1zbS0xMCI+CiAgICAgICAgICA8dGV4dGFyZWEgY2xhc3M9ImZvcm0tY29udHJvbCIgaWQ9ImZlbmdfc2hvd190b2tlbiIgcm93cz0iNCI+ClVEU1dhbjNnUE5sUlJScnFlMWhqUzRFU2gvUjFSaG9ZMXBFdlFIZW1oN2lsSlhXVHFDR21Iblp0bjdGOXBTcmE2MTc2Q1dxTlo1WUZjZm1hdXhLczk4ajNEdGUwNXhUTm5rdkZQV1BYV0Z3TlhKeXcyc2hkdWVHT241azJwWW82eU0zN3pTRzZ2WGRhYzlXVU1rQVcyUVdJTWM4a0lubnd6ck9VME14VGNnYzwvdGV4dGFyZWEKICAgICAgICAgID4KICAgICAgICA8L2Rpdj4KICAgICAgPC9kaXY+CiAgICA8L2Rpdj4KCiAgICA8IS0tIOaPkOS6pOaMiemSriAtLT4KICAgIDxkaXYgY2xhc3M9InRleHQtY2VudGVyIj4KICAgICAgPGJ1dHRvbiBvbmNsaWNrPSJoYW5kbGVBdXRoKCkiIHR5cGU9ImJ1dHRvbiIgY2xhc3M9ImJ0biBidG4tcHJpbWFyeSI+CiAgICAgICAg55m75b2V6I635Y+WIHRva2VuCiAgICAgIDwvYnV0dG9uPgogICAgPC9kaXY+CgogICAgPCEtLSDnu5Pmnpzlj43ppoggLS0+CiAgICA8ZGl2IGNsYXNzPSJyZXN1bHQtd3JhcCI+CiAgICAgIDxkaXYgY2xhc3M9Im1lc3NhZ2UiPuivt+Whq+WFpeiHquW3seeahOWHpOWHsOengOi0puWPt+S/oeaBr+WQjueCueWHu+eZu+W9leiOt+WPliB0b2tlbjwvZGl2PgogICAgICA8ZGl2IGNsYXNzPSJhdXRoLXRva2VuIj4KICAgICAgICA8dGV4dGFyZWEKICAgICAgICAgIGNsYXNzPSJmb3JtLWNvbnRyb2wiCiAgICAgICAgICBkaXNhYmxlZAogICAgICAgICAgaWQ9ImZlbmdfYXV0aF90b2tlbiIKICAgICAgICAgIHBsYWNlaG9sZGVyPSLmjojmnYPkv6Hmga/lnKjov5nph4zlsZXnpLoiCiAgICAgICAgPjwvdGV4dGFyZWE+CiAgICAgIDwvZGl2PgogICAgPC9kaXY+CiAgPC9ib2R5PgoKICA8c2NyaXB0PgogICAgY29uc3QgaGFuZGxlQXV0aCA9ICgpID0+IHsKICAgICAgY29uc3QgcGhvbmUgPSBkb2N1bWVudC5xdWVyeVNlbGVjdG9yKCIjcGhvbmUiKS52YWx1ZTsKICAgICAgY29uc3QgcGFzc3dvcmQgPSBkb2N1bWVudC5xdWVyeVNlbGVjdG9yKCIjcGFzc3dvcmQiKS52YWx1ZTsKICAgICAgY29uc3QgY29kZSA9IGRvY3VtZW50LnF1ZXJ5U2VsZWN0b3IoIiNjb2RlIikudmFsdWU7CiAgICAgIGNvbnN0IGZlbmdTaG93VG9rZW4gPSBkb2N1bWVudC5xdWVyeVNlbGVjdG9yKCIjZmVuZ19zaG93X3Rva2VuIikudmFsdWU7CiAgICAgIGNvbnN0IGRhdGEgPSB7cGhvbmUsIHBhc3N3b3JkLCBjb2RlLCB0aWNrZXQ6IGZlbmdTaG93VG9rZW59OwoKICAgICAgY29uc3QgbWVzc2FnZUVsbSA9IGRvY3VtZW50LnF1ZXJ5U2VsZWN0b3IoIi5yZXN1bHQtd3JhcCAubWVzc2FnZSIpOwogICAgICBtZXNzYWdlRWxtLmlubmVySFRNTCA9ICLmraPlnKjmjojmnYPkuK0sIOivt+eojeWQji4uLiI7CiAgICAgIGNvbnN0IHRva2VuRWxtID0gZG9jdW1lbnQucXVlcnlTZWxlY3RvcigiLnJlc3VsdC13cmFwICNmZW5nX2F1dGhfdG9rZW4iKTsKCiAgICAgIC8vIOWPkei1tyBQT1NUIOivt+axggogICAgICBmZXRjaCgiaHR0cDovL20uZmVuZ3Nob3dzLmNvbS91c2VyL29hdXRoL2xvZ2luIiwgewogICAgICAgIG1ldGhvZDogIlBPU1QiLCAvLyDkvb/nlKggUE9TVCDmlrnms5UKICAgICAgICBoZWFkZXJzOiB7CiAgICAgICAgICAiQ29udGVudC1UeXBlIjogImFwcGxpY2F0aW9uL2pzb24iLCAvLyDmjIflrpror7fmsYLkvZPnsbvlnovkuLogSlNPTgogICAgICAgIH0sCiAgICAgICAgYm9keTogSlNPTi5zdHJpbmdpZnkoZGF0YSksIC8vIOWwhuWvueixoei9rOaNouS4uiBKU09OIOWtl+espuS4sgogICAgICB9KQogICAgICAgIC50aGVuKChyZXNwb25zZSkgPT4gewogICAgICAgICAgaWYgKCFyZXNwb25zZS5vaykgewogICAgICAgICAgICB0aHJvdyBuZXcgRXJyb3IoYEhUVFAgZXJyb3IhIFN0YXR1czogJHtyZXNwb25zZS5zdGF0dXN9YCk7CiAgICAgICAgICB9CiAgICAgICAgICByZXR1cm4gcmVzcG9uc2UuanNvbigpOyAvLyDop6PmnpAgSlNPTiDlk43lupTkvZMKICAgICAgICB9KQogICAgICAgIC50aGVuKChkYXRhKSA9PiB7CiAgICAgICAgICBpZiAoZGF0YS5zdGF0dXMgIT09ICcwJykgewogICAgICAgICAgICB0aHJvdyBuZXcgRXJyb3IoYOaOiOadg+Wksei0pTogJHtkYXRhLm1lc3NhZ2V9YCk7CiAgICAgICAgICB9CiAgICAgICAgICBjb25zdCB0b2tlbiA9IGRhdGE/LmRhdGE/LnRva2VuOwogICAgICAgICAgbWVzc2FnZUVsbS5pbm5lckhUTUwgPSAn5o6I5p2D5oiQ5Yqf77yM6K+35bCG6I635Y+W5Yiw55qE5YC86K6+572u5Yiw56iL5bqP55qE546v5aKD5Y+Y6YeP5Lit77yMa2V5OiBmZW5nX3Rva2VuJzsKICAgICAgICAgIHRva2VuRWxtLnZhbHVlID0gdG9rZW47CiAgICAgICAgfSkKICAgICAgICAuY2F0Y2goKGVycm9yKSA9PiB7CiAgICAgICAgICBtZXNzYWdlRWxtLmlubmVySFRNTCA9IGVycm9yOwogICAgICAgICAgc2V0VGltZW91dCgoKSA9PiBtZXNzYWdlRWxtLmlubmVySFRNTCA9ICfor7floavlhaXoh6rlt7HnmoTlh6Tlh7Dnp4DotKblj7fkv6Hmga/lkI7ngrnlh7vnmbvlvZXojrflj5YgdG9rZW4nLCA1MDAwKTsKICAgICAgICB9KTsKICAgIH07CiAgPC9zY3JpcHQ+CjwvaHRtbD4K`

	// ConfigPageHtml 配置页地址
	ConfigPageHtml = `PCFET0NUWVBFIGh0bWw+CjxodG1sPgogIDxoZWFkPgogICAgPG1ldGEgY2hhcnNldD0iVVRGLTgiIC8+CiAgICA8bWV0YSBuYW1lPSJ2aWV3cG9ydCIgY29udGVudD0id2lkdGg9ZGV2aWNlLXdpZHRoLCBpbml0aWFsLXNjYWxlPTEuMCIgLz4KICAgIDxsaW5rIHJlbD0iaWNvbiIgdHlwZT0iaW1hZ2UveC1pY29uIiBocmVmPSJkYXRhOmltYWdlL3ZuZC5taWNyb3NvZnQuaWNvbjtiYXNlNjQsQUFBQkFBRUFJQ0FBQUFFQUlBQ29FQUFBRmdBQUFDZ0FBQUFnQUFBQVFBQUFBQUVBSUFBQUFBQUFBQkFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUQ4L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy92OS9mLy8vLy8vLy8vLy8vLy8vLy82L1B6LysvejgvLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vNy9mei8rLzM5Ly8vLy8vLy8vLy8vLy8vLy8vdjkvZi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi83L2YzLy8vLy8vK1BrNVArUWxKWC94TVhFLy8vLy8vLy8vLy8vN2UvMi85dmQ1UC9OenRUL3hNWEsvOEhDeHYvQnc4Zi94Y2ZNLzh6TzAvL1oyK0gvN083MS8vLy8vLy8vLy8vL3ZzQy8vNU9ZbWYvbDV1Yi8vLy8vLy92OS9mLzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy9uNysvLy8vLy8vbVpxYS8wcDFrLzlXZjZML2NYQnYvNCtQZnYrRWdsNy9pb2hXLzVDT1V2K1VrbEgvbUpaUy81aVZVditWazFML2tJMVMvNG1IVlArRmhGMy9rWkY3LzIxdGFmOVZmNlAvU25XVC81K2hvUC8vLy8vLytmdjcvL3orL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3LysvMzkvLy8vLy8vMDlmYi9XbWR2LzM2RllQKzBzRnYvMHM1cC8rUGZkZi9vNVhyLzZ1ZDgvK3JuZmYvcTUzMy82dWQ5LytybmZmL3E1M3ovNmVWNi8rUGZkZi9VMEd2L3RMQmMvNHFSWmY5YVpXbi83dS95Ly8vLy8vLzcvZjMvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzYvUHovLy8vLy81NmZuZitrbmxQLzN0bHgvN200YWYrbG8yRC9sNU5ZLzQySlV2K0dnVTcvZzMxTS80RjhTLytBZWtyL2dYeE0vNGVEVC8rUWpWWC9tNXBiLzdDdVpQL2IySGIvNU41MS82dWxWZitQa0lqLy8vLy8vL3Y5L1AvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3LysvejgvLy8vLy8vUzFOdi9pSVpQLzgvTWJ2OVBURDcvT0M0dS96TXRNUDh5TWpmL01EVSsvekk2US84eVBVYi9NVHhHL3pBOFJ2OHpQVWIvTVRjLy96SXhPUDgyTGpEL09UQXovMlJpUmYvT3kyLy81T0Y3LzVHT1RQKyt2OFQvLy8vLy8vcjgvUC84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzUrL3YvLy8vLy81cWJsLy9DdjJIL2lJWlYvekV2T2Y5QmZZNy9UNjdJLzFpLzJ2OVR3ZUQvVHIvaC8xdk83UDlhemV6L1RzTGwvMVhJNlA5YXlPVC9TckRRLzB5cHcvOUhhbkwvTnk4MC82cXBZUC9lMjNuL3hjRmsvNDZOZnYvLy8vLy8rdno4Ly96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3Y5L1AvLy8vLy9pNHQzLzlQUmEvOXVaa1QvT0U1Wi8xcmIvdjlSemZiL1c5WDUvMS9aK3Y5V3ovVC9VY253LzE3VytQOWMxZmIvVU1qdi8xblM5djllMS9uL1ZOZi8vMDJ3eS84Mkx6UC9sNVZaLzkzWmR2L1YwVy8vaVlkaS8rL3grUC8vLy8vLy9QMysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL3YvLy8vWDMvZitKaUduLzJkZHcvMTlWTy84NlhXei9YOXY4LzFuTjd2OU53K3IvV2MveC8xN1Y5ZjlWek8vL1VNZnMvMTNUOVA5YjB2UC9UOFhyLzFmTjcvOWUxdmovUzdQVC96TTBPLytMaUZQLzNObDIvOW5XYy8rTWkxci80T0xwLy8vLy8vLzcvZjMvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3Y5L2YvLy8vLy83TzcxLzRpSFlmL2EyWEwvV1U0NC96ZGdjdjlVMC9uL1h0UHovMW5ROHY5TnhPdi9XTkR5LzE3Vjl2OVV6UEQvVU1mdC8xM1U5ZjljMC9UL1Q4WHIvMW5TOXY5WngrVC9NemhCLzROL1R2L2IySFgvMnRaMC80K05WLy9XMk4vLy8vLy8vL3I4L1AvOC92Ny8vUDcrLy96Ky92LzgvdjcvL2YvLy8vLy8vLy8wOVBuL2pZeGsvOXJZY3Y5VlNqYi9QMnA1LzFUVSt2OVN4K3ovWHRYMi8xblE4LzlNeE92L1dNL3kvMTdXOXY5VnpQRC9VTWZ0LzEzVTlmOWIwdlAvVU1udy8xUEQ0LzgyUFVYL2YzcE0vK0xmZVAvaDNYYi9scFJaLzl6YzRmLy8vLy8vL1A3Ly8veisvdi84L3Y3Ly9QNysvL3orL3YvLy8vLy9rNXFkLzNxRmt2K0ZnbGovM05wMC8xVktOdjg5Wm5iL1lOLy8vMUhIN1A5U3llNy9YdFgyLzFuUjgvOU54T3YvVjgveS8xN1c5djlWelBEL1VNZnQvMXpUOVA5ZDF2ai9UTDNlL3pJNVEvK0VmMC8vaElKUi8yVmpSZitSamxUL2RuNkkvNVNibnYvNStmbi8vZi8vLy96Ky92LzcvZjMvLy8vLy8rVGk0UDlNWm5YL1phTFQvM2w5V2YvZjIzVC9XRTQ1L3pkY2JQOVUwZmYvWHRUei8xTEo3djlTeWU3L1h0WDIvMXJSOC85TnhPdi9WODd4LzE3Vzl2OVV6UEQvVU1icy8xN1grZjlXdzkvL016ZEEvNFNBVC85eGIwci9VRTAvLzRhRVRmOWptOGIvVEd1QS85blgxZi8vLy8vLysvejkvL3orL3YvNy9mNy8vLy8vLzdyQXd2OXBncGovZVh4WS8rSGVkZjlkVkR2L08xbGwvMVRUK2Y5VHlPMy9YdFQxLzFMSTdmOVF4dXYvWGRQei8xclE4ZjlOdytuL1Y4M3cvMS9WOWY5VnkrNy9WTkg1LzFXKzJ2ODBNVGovaVlaVC83aTFaUCtrb1ZyL21KZGEvMkYxaWYrOXc4WC8vLy8vLy92OS92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly8vLy8vK25uNnYrRGdWdi80dDkxLzJoalJmODVRRWovVmI3WC8xUE84LzlUMFBYL1g5MysvMVRTK1A5VTB2ci9ZTjcvLzE3Yy8vOU95L1AvVnRMMi8xelU5UDlTeU92L1FvU1Yvend6T1ArSWgxSC9WVk5CL3pjME92OXhiMFgvMnRqWi8vLy8vLy83L2YzLy9QNysvL3orL3YvOC92Ny8vUDcrLy96OS92LysvLzcvN3ZINC80cUpZLy9pM25YL2taQlcvekVzTmY4M09rSC9PRTlhL3paVVkvODdYV3YvUW1oMS96TmJhUDh3V0dQL05sdGwvMEJnYlA4NFUyRC9PMDlaL3poQVIvODdOVHYvTmpNMy83S3ZZZitQakZYL2VIVk4vNG1IVlAvWDJ1SC8vLy8vLy92OC9mLzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvLzcvLy8vMitQNy9mWHhkLzlmVGEvL1kxSFgvbUpaWS8zTnZTZjlsWFQ3L1hGTTYvMXBSUFA4K01SNy9nM2R3LzZTYWxmK1Nob0wvUFRBZy8xcFNQZjlhVWp6L1pGOUQvMjV0U2YrZ25sci8xOVJ6LytEZGQvL2gzVy8vY25CSi8rTGs2Ly8vLy8vLysvMzkvL3orL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvK3Z6OC8vLy8vLys0dWJqL2ZIdFMvOUhOY1AvZDJuZi80TjE0LytYaWV2L1oxM0QvMU5CcS8ydDVaUDk3bWEvL2hKV2UvNEtlc2Y5ZmNtNy94OFJuLzkzYmNmL2o0SGYvNE4xNC85cldkZi9aMVhUL29xQmQvM2QyWC8rd3NiRC8vLy8vLy92OS9mLzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8rZnY2Ly8vLy8vK1ZscGIvczdCYS85dlllZi9KeG1uL2pveFYvNUdSY3Y5NGMwNy9VSEtKLzNldTF2OUpZRy8vY2FQRi8xcUdxUDlyWjBUL2xaUnYvNCtOVnYvSXhXai8xTkYwLzlmVWIvOTlmRnovNnV6NC8vLy8vLy83L2YzLy9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi81Ky92Ly8vLy8vNnl0clArcXAxai8yOWR6LzRHQVlQL016dGYvK1ByLy8vLy8vLytvcXFyL0tTa2svMEEzTS84bkpTTC9scGlZLy9yOC8vLzE5LzMvME5MYi80S0JaUC9QeTJ6LzA4OXQvNGVIYS8vNi9QLy8rL3o3Ly96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy9yOC9QLy8vLy8veGNmTS82T2hXLytycUZ2L3VidkMvLy8vLy8vLy8vLy92cis5LzdxN3Z2K09qWG4vcGFOUi80bUljdi9GeHNyL3VMcTQvLy8vLy8vLy8vLy94TVhOLzUyYldmL1J6V24va1pDQS8vLy8vLy83L1B6Ly9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Lyt2ejgvLy8vLy8vaTVQSC9oWVJZLzUyYlcvL1oyK0wvLy8vLy8vLy8vLzgrUGp6L1NVbEwvNmlvbnYrMHNGSC9yNitmLzBwS1RmODBOREwvK1ByNi8vLy8vLy9qNWV6L25KcGsvN0N0Vi8rUWtaai8vLy8vLy9uNysvLzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy92OS9mLy8vLy8vNnV6dy80U0ViLzkzZFVqL2w1UlIvNzYvdy8vLy8vLy8vLy8vLzlyYjJmL2EzT0gva1pGKy84ZkRYditRajNYLzFkYmQvOUhTei8vKy8vLy8vLy8vLzhMRHpQK25wRnYvall0Uy8yOXRTZisxdHJQLy8vLy8vL3I4L1AvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Lyt2djcvLy8vLy8rWm1wWC90N05YLzRKL1R2OXBaejcvbDVaci85WFczLy81Ky8vLy9mLy8vNmVvclArWmwxWC81T0Y3LzZTaFYvK2FtNXIvK2Z2Ly8vZjUvLy9VMXVIL21KZHcvNW1XVS85TVNqYi8wczVyLzRTQ1dmL3E3UFAvLy8vLy8vejkvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzYrL3YvLy8vLy81eWRsditxcGxIL3RiSmsvNDJLVGY5ZFd6UC9lM2xGLzVXVWNQK1VrMm4vcmFwYi8rSGRkLy9YMDNULzRkNTQvN1N4WGYrVGtXVC9sNVp3LzRlRlRQOTZkei8vYkdvLy82MnFYLysrdW1ML2g0VmMvK3Z0OVAvLy8vLy8vUDMrLy96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3Y5L2YvLy8vLy82dXp3LzRXRWJ2K1lsVmYvaTRwcy83dTh4UCthbW96L2YzMVQvNUtQVHYrbW8xZi90YkpsLzd1NFovKzV0bWYvcnF0Zi81YVRUditGZzFIL2lvcHgvN3E3d2YrYm5KTC9sWkpXLzRTQ1ZQKzB0YmIvLy8vLy8vcjgvUC84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL2Y3Ky8vdjkvZi8vLy8vLyt2ei8vOVhYMy8vMDl2di8vLy8vLy8vLy8vL3E3UFQvMGRMYS82V25xLzhvSmliL05USXgveVFoSVA5TlRWRC95ODNVLytEaTZ2LysvLy8vLy8vLy8vLy8vLy9XMk9ELzR1VHIvLy8vLy8vNy9mMy8vZjcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vZjcrLy92OS9mLzkvLzcvLy8vLy8vNy8vLy81Ky92Lyt2djcvLzcvL3YvLy8vLy9tNXVhLzRhSGh2Ky92Ny8vcksydC8ydHJhdi8vLy8vLy8vLy8vL3I3Ky8vNSsvdi8rLzM4Ly8vLy8vLy8vLy8vK3Z6OC8vMysvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvNy9mMy8vUDcrLy96Ky92LzgvdjcvLy8vLy81U1ZsUCtNakl6Ly8vLy8vLy8vLy8vLy8vLy93OFRFLzI5dmJ2Ly8vLy8vLy8vLy8vdjkvZi84L3Y3LysvMzkvL3Y5L2YvOC92Ny8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9mNysvLzMrL3YvOC92Ny8vUDcrLy8zKy92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QMysvLy8vLy8rZm41Ly9uWjJkLy8vLy8vLzYvUHovK2Z2Ny8vYjQrUC8vLy8vL3VicTUvNGlJaVAvLy8vLy8vUDcrLy96Ky92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vZjcrLy96Ky92LzkvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvL3orL3YvOC92Ny8vUDcrLy9yOC9QLy8vLy8vd2NMQy81aVltUC8vLy8vLytmdjcvL3orL3YvOS92Ny8vZjcvLy9qNSt2Ly8vLy8vb0tHaC84Ykh4Ly8vLy8vLyt2ejgvL3orL3YvOC92Ny8vUDcrLy8zKy92LzkvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvLzMrL3YvOC92Ny8vUDcrLy8zKy92LzgvdjcvL1A3Ky8veisvdi84L3Y3Ly9QNysvLzMrLy8vNy9QMy8vLy8vLy92OS9mLzgvdjcvL2Y3Ky8vMysvdi85L3Y3Ly9mNysvL3Y4L1AvLy8vLy8vdi8vLy96Ky92LzgvdjcvL1A3Ky8vMysvdi84L3Y3Ly9QNysvLzMrL3YvOS92Ny8vUDcrLy8zKy92LzkvdjcvQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBPSI+CiAgICA8dGl0bGU+bGl2ZS1zZXJ2ZXIg6YWN572u6aG1PC90aXRsZT4KICAgIDwhLS0g5byV5YWlIEJvb3RzdHJhcCBDU1MgLS0+CiAgICA8bGluawogICAgICBocmVmPSJodHRwczovL2Zhc3RseS5qc2RlbGl2ci5uZXQvbnBtL2Jvb3RzdHJhcEA1LjMuMC9kaXN0L2Nzcy9ib290c3RyYXAubWluLmNzcyIKICAgICAgcmVsPSJzdHlsZXNoZWV0IgogICAgLz4KICAgIDxzdHlsZT4KICAgICAgLnBhZ2UtY29udGFpbmVyIHsKICAgICAgICBwYWRkaW5nOiAyMHB4IDEwcHg7CiAgICAgIH0KICAgICAgLnBhZ2UtY29udGFpbmVyIGlucHV0LAogICAgICAucGFnZS1jb250YWluZXIgdGV4dGFyZWEsCiAgICAgIC5wYWdlLWNvbnRhaW5lciAuYnV0dG9uLWdyb3VwIHsKICAgICAgICBtYXJnaW46IDhweCAwOwogICAgICB9CiAgICAgIC5lbnYtd3JhcCAuZW52LXZhbHVlIHsKICAgICAgICBvdmVyZmxvdy13cmFwOiBicmVhay13b3JkOwogICAgICAgIHdvcmQtYnJlYWs6IGJyZWFrLWFsbDsKICAgICAgfQogICAgPC9zdHlsZT4KICA8L2hlYWQ+CiAgPGJvZHk+CiAgICA8ZGl2IGNsYXNzPSJwYWdlLWNvbnRhaW5lciI+CiAgICAgIDxoMj5saXZlLXNlcnZlciDphY3nva7pobU8L2gyPgoKICAgICAgPCEtLSDlsZXnpLrnlKjmiLfkvKDpgJLnmoTnqIvluo/lr4bpkqUgLS0+CiAgICAgIDxkaXYgY2xhc3M9Im1iLTMgcm93IHNlY3JldC13cmFwIj4KICAgICAgICA8bGFiZWwgY2xhc3M9ImNvbC1zbS0yIGNvbC1mb3JtLWxhYmVsIj7nqIvluo/lr4bpkqXvvJo8L2xhYmVsPgogICAgICAgIDxkaXYgY2xhc3M9ImNvbC1zbS0xMCI+CiAgICAgICAgICA8aW5wdXQgY2xhc3M9ImZvcm0tY29udHJvbCBzZWNyZXQiIHR5cGU9InRleHQiIGRpc2FibGVkIC8+CiAgICAgICAgPC9kaXY+CiAgICAgIDwvZGl2PgoKICAgICAgPCEtLSDnjq/looPlj5jph48gLS0+CiAgICAgIDxkaXYgY2xhc3M9ImVudi13cmFwIj4KICAgICAgICA8aDM+546v5aKD5Y+Y6YePPC9oMz4KICAgICAgICA8ZGl2IGNsYXNzPSJtYi0zIHJvdyI+CiAgICAgICAgICA8ZGl2IGNsYXNzPSJjb2wtc20tNiI+CiAgICAgICAgICAgIDxpbnB1dAogICAgICAgICAgICAgIGNsYXNzPSJmb3JtLWNvbnRyb2wgZW52LWtleSIKICAgICAgICAgICAgICB0eXBlPSJ0ZXh0IgogICAgICAgICAgICAgIHBsYWNlaG9sZGVyPSLlnKjmraTovpPlhaUga2V5LCDlpoI6IGZlbmdfdG9rZW4iCiAgICAgICAgICAgIC8+CiAgICAgICAgICA8L2Rpdj4KICAgICAgICAgIDxkaXYgY2xhc3M9ImNvbC1zbS02IGJ1dHRvbi1ncm91cCI+CiAgICAgICAgICAgIDxidXR0b24gdHlwZT0iYnV0dG9uIiBjbGFzcz0iYnRuIGJ0bi1zdWNjZXNzIiBvbmNsaWNrPSJnZXRFbnYoKSI+CiAgICAgICAgICAgICAg5p+lIOivogogICAgICAgICAgICA8L2J1dHRvbj4KICAgICAgICAgICAgPGJ1dHRvbiB0eXBlPSJidXR0b24iIGNsYXNzPSJidG4gYnRuLXByaW1hcnkiIG9uY2xpY2s9InNldEVudigpIj4KICAgICAgICAgICAgICDmt7sg5YqgCiAgICAgICAgICAgIDwvYnV0dG9uPgogICAgICAgICAgICA8YnV0dG9uIHR5cGU9ImJ1dHRvbiIgY2xhc3M9ImJ0biBidG4tZGFuZ2VyIiBvbmNsaWNrPSJkZWxFbnYoKSI+CiAgICAgICAgICAgICAg5YigIOmZpAogICAgICAgICAgICA8L2J1dHRvbj4KICAgICAgICAgIDwvZGl2PgogICAgICAgICAgPGRpdiBjbGFzcz0iY29sLXNtLTEyIj4KICAgICAgICAgICAgPHRleHRhcmVhCiAgICAgICAgICAgICAgY2xhc3M9ImZvcm0tY29udHJvbCBlbnYtdmFsdWUiCiAgICAgICAgICAgICAgcm93cz0iNiIKICAgICAgICAgICAgICBzdHlsZT0ib3ZlcmZsb3c6IGhpZGRlbjsgcmVzaXplOiBub25lIgogICAgICAgICAgICAgIHBsYWNlaG9sZGVyPSLlnKjmraTovpPlhaXmg7Pmm7TmlrDnmoQgdmFsdWUsIOmAmui/h+afpeivouW+l+WIsOeahOWAvOS5n+S8muaYvuekuuWcqOatpCIKICAgICAgICAgICAgPjwvdGV4dGFyZWE+CiAgICAgICAgICA8L2Rpdj4KICAgICAgICA8L2Rpdj4KICAgICAgPC9kaXY+CgogICAgICA8IS0tIOm7keWQjeWNlSAtLT4KICAgICAgPGRpdiBjbGFzcz0iYmxhY2staXAtd3JhcCI+CiAgICAgICAgPGgzPum7keWQjeWNlTwvaDM+CiAgICAgICAgPGRpdiBjbGFzcz0ibWItMyByb3ciPgogICAgICAgICAgPGRpdiBjbGFzcz0iY29sLXNtLTEwIj4KICAgICAgICAgICAgPGlucHV0CiAgICAgICAgICAgICAgY2xhc3M9ImZvcm0tY29udHJvbCBibGFjay1pcCIKICAgICAgICAgICAgICB0eXBlPSJ0ZXh0IgogICAgICAgICAgICAgIHBsYWNlaG9sZGVyPSLor7fovpPlhaXopoHliqDlhaXpu5HlkI3ljZXnmoQgaXAsIOWmgjogMS4xLjEuMSIKICAgICAgICAgICAgLz4KICAgICAgICAgIDwvZGl2PgogICAgICAgICAgPGRpdiBjbGFzcz0iY29sLXNtLTIgYnV0dG9uLWdyb3VwIj4KICAgICAgICAgICAgPGJ1dHRvbgogICAgICAgICAgICAgIHR5cGU9ImJ1dHRvbiIKICAgICAgICAgICAgICBjbGFzcz0iYnRuIGJ0bi1wcmltYXJ5IgogICAgICAgICAgICAgIG9uY2xpY2s9InNldEJsYWNrSXAoKSIKICAgICAgICAgICAgPgogICAgICAgICAgICAgIOa3uyDliqAKICAgICAgICAgICAgPC9idXR0b24+CiAgICAgICAgICA8L2Rpdj4KICAgICAgICA8L2Rpdj4KICAgICAgPC9kaXY+CgogICAgICA8IS0tIOWcsOWfn+eZveWQjeWNlSAtLT4KICAgICAgPGRpdiBjbGFzcz0id2hpdGUtYXJlYS13cmFwIj4KICAgICAgICA8aDM+5Zyw5Z+f55m95ZCN5Y2VPC9oMz4KICAgICAgICA8ZGl2IGNsYXNzPSJtYi0zIHJvdyI+CiAgICAgICAgICA8ZGl2IGNsYXNzPSJjb2wtc20tOCI+CiAgICAgICAgICAgIDxpbnB1dAogICAgICAgICAgICAgIGNsYXNzPSJmb3JtLWNvbnRyb2wgd2hpdGUtYXJlYSIKICAgICAgICAgICAgICB0eXBlPSJ0ZXh0IgogICAgICAgICAgICAgIHBsYWNlaG9sZGVyPSLor7fovpPlhaXopoHliqDlhaXnmb3lkI3ljZXnmoTlnLDln58sIOWmgjog5bm/5LicL+S9m+WxsS/ljZfmtbciCiAgICAgICAgICAgIC8+CiAgICAgICAgICA8L2Rpdj4KICAgICAgICAgIDxkaXYgY2xhc3M9ImNvbC1zbS00IGJ1dHRvbi1ncm91cCI+CiAgICAgICAgICAgIDxidXR0b24KICAgICAgICAgICAgICB0eXBlPSJidXR0b24iCiAgICAgICAgICAgICAgY2xhc3M9ImJ0biBidG4tcHJpbWFyeSIKICAgICAgICAgICAgICBvbmNsaWNrPSJzZXRXaGl0ZUFyZWEoJ3NldCcpIgogICAgICAgICAgICA+CiAgICAgICAgICAgICAg5re7IOWKoAogICAgICAgICAgICA8L2J1dHRvbj4KICAgICAgICAgICAgPGJ1dHRvbgogICAgICAgICAgICAgIHR5cGU9ImJ1dHRvbiIKICAgICAgICAgICAgICBjbGFzcz0iYnRuIGJ0bi1kYW5nZXIiCiAgICAgICAgICAgICAgb25jbGljaz0ic2V0V2hpdGVBcmVhKCdkZWwnKSIKICAgICAgICAgICAgPgogICAgICAgICAgICAgIOenuyDpmaQKICAgICAgICAgICAgPC9idXR0b24+CiAgICAgICAgICA8L2Rpdj4KICAgICAgICA8L2Rpdj4KICAgICAgPC9kaXY+CiAgICA8L2Rpdj4KICA8L2JvZHk+CgogIDxzY3JpcHQ+CiAgICAvLyDojrflj5bln7rmnKzlj4LmlbAKICAgIGNvbnN0IGdldEJhc2VQYXJhbXMgPSAoKSA9PiB7CiAgICAgIGNvbnN0IHNlY3JldCA9IG5ldyBVUkwod2luZG93LmxvY2F0aW9uLmhyZWYpLnNlYXJjaFBhcmFtcy5nZXQoInNlY3JldCIpOwogICAgICByZXR1cm4gewogICAgICAgIHNlY3JldCwgLy8g56iL5bqP5a+G6ZKlCiAgICAgICAgaG9zdDogIiIsIC8vIOWPkemAgeivt+axgueahOS4u+acuuWcsOWdgAogICAgICB9OwogICAgfTsKCiAgICAvLyDmoLnmja7mlofmnKzmoYblhoXlrrnoh6rliqjosIPmlbTnjq/looPlj5jph4/moYbnmoTpq5jluqYKICAgIGNvbnN0IGF1dG9SZXNpemVFbnZWYWx1ZVRleHRhcmVhID0gKGluaXRFdmVudCA9IGZhbHNlKSA9PiB7CiAgICAgIGNvbnN0IHRleHRhcmVhID0gZG9jdW1lbnQucXVlcnlTZWxlY3RvcigiLmVudi13cmFwIC5lbnYtdmFsdWUiKTsKCiAgICAgIC8vIOaguOW/g+iwg+aVtOWHveaVsAogICAgICBjb25zdCBhdXRvUmVzaXplID0gKCkgPT4gewogICAgICAgIHRleHRhcmVhLnN0eWxlLmhlaWdodCA9ICJhdXRvIjsKICAgICAgICB0ZXh0YXJlYS5zdHlsZS5oZWlnaHQgPSBgJHt0ZXh0YXJlYS5zY3JvbGxIZWlnaHR9cHhgOwogICAgICB9OwoKICAgICAgaWYgKGluaXRFdmVudCkgewogICAgICAgIC8vIOi+k+WFpeS6i+S7tuebkeWQrAogICAgICAgIHRleHRhcmVhLmFkZEV2ZW50TGlzdGVuZXIoImlucHV0IiwgYXV0b1Jlc2l6ZSk7CgogICAgICAgIC8vIOeql+WPo+WwuuWvuOWPmOWMluebkeWQrO+8iOW4pumYsuaKlu+8iQogICAgICAgIGxldCByZXNpemVUaW1lcjsKICAgICAgICB3aW5kb3cuYWRkRXZlbnRMaXN0ZW5lcigicmVzaXplIiwgKCkgPT4gewogICAgICAgICAgY2xlYXJUaW1lb3V0KHJlc2l6ZVRpbWVyKTsKICAgICAgICAgIHJlc2l6ZVRpbWVyID0gc2V0VGltZW91dChhdXRvUmVzaXplLCAxNTApOwogICAgICAgIH0pOwogICAgICB9CgogICAgICAvLyDliJ3lp4vljJbmiafooYwKICAgICAgYXV0b1Jlc2l6ZSgpOwogICAgfTsKCiAgICB3aW5kb3cub25sb2FkID0gKCkgPT4gewogICAgICAvLyDorr7nva7nqIvluo/lr4bpkqUKICAgICAgY29uc3QgeyBzZWNyZXQgfSA9IGdldEJhc2VQYXJhbXMoKTsKICAgICAgZG9jdW1lbnQucXVlcnlTZWxlY3RvcigiLnBhZ2UtY29udGFpbmVyIC5zZWNyZXQiKS52YWx1ZSA9IHNlY3JldDsKCiAgICAgIGF1dG9SZXNpemVFbnZWYWx1ZVRleHRhcmVhKHRydWUpOwogICAgfTsKCiAgICAvLyDojrflj5bovpPlhaXmoYbnmoTlgLwKICAgIGNvbnN0IGdldElucHV0VmFsdWUgPSAoc2VsZWN0b3IgPSAiIikgPT4gewogICAgICByZXR1cm4gKGRvY3VtZW50LnF1ZXJ5U2VsZWN0b3Ioc2VsZWN0b3IpLnZhbHVlIHx8ICIiKS50cmltKCk7CiAgICB9OwoKICAgIC8vIOiuvue9rui+k+WFpeahhueahOWAvAogICAgY29uc3Qgc2V0SW5wdXRWYWx1ZSA9IChzZWxlY3RvciA9ICIiLCB2YWx1ZSA9ICIiKSA9PiB7CiAgICAgIGNvbnN0IHRhcmdldCA9IGRvY3VtZW50LnF1ZXJ5U2VsZWN0b3Ioc2VsZWN0b3IpOwogICAgICBpZiAoIXRhcmdldCkgewogICAgICAgIHJldHVybjsKICAgICAgfQogICAgICB0YXJnZXQudmFsdWUgPSB2YWx1ZTsKICAgIH07CgogICAgLy8gR0VUIOivt+axguaMh+WumuWcsOWdgAogICAgY29uc3QgZmV0Y2hVcmwgPSAoCiAgICAgIHVybCA9ICIiLAogICAgICBzdWNjZXNzRnVuYyA9ICgpID0+IHt9LAogICAgICBlcnJvckZ1bmMgPSAoZXJyKSA9PiB7fQogICAgKSA9PiB7CiAgICAgIGZldGNoKHVybCkKICAgICAgICAudGhlbigocmVzKSA9PiB7CiAgICAgICAgICBpZiAocmVzLm9rKSB7CiAgICAgICAgICAgIHJldHVybiBzdWNjZXNzRnVuYygpOwogICAgICAgICAgfQogICAgICAgICAgZXJyb3JGdW5jKHJlcy5zdGF0dXNUZXh0KTsKICAgICAgICB9KQogICAgICAgIC5jYXRjaCgoZXJyKSA9PiB7CiAgICAgICAgICBlcnJvckZ1bmMoZXJyKTsKICAgICAgICB9KTsKICAgIH07CgogICAgLy8g5p+l6K+i546v5aKD5Y+Y6YePCiAgICBjb25zdCBnZXRFbnYgPSAoKSA9PiB7CiAgICAgIGNvbnN0IGVudktleSA9IGdldElucHV0VmFsdWUoIi5lbnYtd3JhcCAuZW52LWtleSIpOwogICAgICBpZiAoZW52S2V5ID09PSAiIikgewogICAgICAgIHJldHVybiBhbGVydCgi6L6T5YWl5Y+C5pWw5LiN6IO95Li656m6Iik7CiAgICAgIH0KCiAgICAgIGNvbnN0IHsgc2VjcmV0LCBob3N0IH0gPSBnZXRCYXNlUGFyYW1zKCk7CiAgICAgIGNvbnN0IHVybCA9IGAke2hvc3R9L2Vudj9rZXk9JHtlbnZLZXl9JnNlY3JldD0ke3NlY3JldH1gOwogICAgICBmZXRjaCh1cmwpCiAgICAgICAgLnRoZW4oKHJlcykgPT4gewogICAgICAgICAgaWYgKCFyZXMub2spIHsKICAgICAgICAgICAgcmV0dXJuIGFsZXJ0KGDmn6Xor6LlpLHotKU6ICR7cmVzLnN0YXR1c1RleHR9YCk7CiAgICAgICAgICB9CiAgICAgICAgICByZXR1cm4gcmVzLnRleHQoKTsKICAgICAgICB9KQogICAgICAgIC50aGVuKChyZXNwVGV4dCkgPT4gewogICAgICAgICAgc2V0SW5wdXRWYWx1ZSgiLmVudi13cmFwIC5lbnYtdmFsdWUiLCByZXNwVGV4dCk7CiAgICAgICAgICBhdXRvUmVzaXplRW52VmFsdWVUZXh0YXJlYShmYWxzZSk7CiAgICAgICAgfSkKICAgICAgICAuY2F0Y2goKGVycikgPT4gewogICAgICAgICAgYWxlcnQoYOafpeivouWksei0pTogJHtlcnJ9YCk7CiAgICAgICAgfSk7CiAgICB9OwoKICAgIC8vIOiuvue9rueOr+Wig+WPmOmHjwogICAgY29uc3Qgc2V0RW52ID0gKCkgPT4gewogICAgICBjb25zdCBlbnZLZXkgPSBnZXRJbnB1dFZhbHVlKCIuZW52LXdyYXAgLmVudi1rZXkiKTsKICAgICAgY29uc3QgZW52VmFsdWUgPSBnZXRJbnB1dFZhbHVlKCIuZW52LXdyYXAgLmVudi12YWx1ZSIpOwogICAgICBpZiAoZW52S2V5ID09PSAiIiB8fCBlbnZWYWx1ZSA9PT0gIiIpIHsKICAgICAgICByZXR1cm4gYWxlcnQoIui+k+WFpeWPguaVsOS4jeiDveS4uuepuiIpOwogICAgICB9CgogICAgICBjb25zdCB7IHNlY3JldCwgaG9zdCB9ID0gZ2V0QmFzZVBhcmFtcygpOwogICAgICBjb25zdCB1cmwgPSBgJHtob3N0fS9lbnY/c2VjcmV0PSR7c2VjcmV0fWA7CiAgICAgIGNvbnN0IGZvcm1EYXRhID0gbmV3IEZvcm1EYXRhKCk7CiAgICAgIGZvcm1EYXRhLmFwcGVuZCgia2V5IiwgZW52S2V5KTsKICAgICAgZm9ybURhdGEuYXBwZW5kKCJ2YWx1ZSIsIGVudlZhbHVlKTsKCiAgICAgIGZldGNoKHVybCwgewogICAgICAgIG1ldGhvZDogIlBPU1QiLAogICAgICAgIGJvZHk6IGZvcm1EYXRhLAogICAgICB9KQogICAgICAgIC50aGVuKChyZXMpID0+IHsKICAgICAgICAgIGlmIChyZXMub2spIHsKICAgICAgICAgICAgcmV0dXJuIGFsZXJ0KCLmt7vliqDmiJDlip8iKTsKICAgICAgICAgIH0KICAgICAgICAgIGFsZXJ0KGDmt7vliqDlpLHotKU6ICR7cmVzLnN0YXR1c1RleHR9YCk7CiAgICAgICAgfSkKICAgICAgICAuY2F0Y2goKGVycikgPT4gewogICAgICAgICAgYWxlcnQoYOa3u+WKoOWksei0pTogJHtlcnJ9YCk7CiAgICAgICAgfSk7CiAgICB9OwoKICAgIC8vIOWIoOmZpOeOr+Wig+WPmOmHjwogICAgY29uc3QgZGVsRW52ID0gKCkgPT4gewogICAgICBjb25zdCBlbnZLZXkgPSBnZXRJbnB1dFZhbHVlKCIuZW52LXdyYXAgLmVudi1rZXkiKTsKICAgICAgaWYgKGVudktleSA9PT0gIiIpIHsKICAgICAgICByZXR1cm4gYWxlcnQoIui+k+WFpeWPguaVsOS4jeiDveS4uuepuiIpOwogICAgICB9CgogICAgICBjb25zdCB7IHNlY3JldCwgaG9zdCB9ID0gZ2V0QmFzZVBhcmFtcygpOwogICAgICBjb25zdCB1cmwgPSBgJHtob3N0fS9lbnY/c2VjcmV0PSR7c2VjcmV0fWA7CiAgICAgIGNvbnN0IGZvcm1EYXRhID0gbmV3IEZvcm1EYXRhKCk7CiAgICAgIGZvcm1EYXRhLmFwcGVuZCgia2V5IiwgZW52S2V5KTsKCiAgICAgIGZldGNoKHVybCwgewogICAgICAgIG1ldGhvZDogIkRFTEVURSIsCiAgICAgICAgYm9keTogZm9ybURhdGEsCiAgICAgIH0pCiAgICAgICAgLnRoZW4oKHJlcykgPT4gewogICAgICAgICAgaWYgKHJlcy5vaykgewogICAgICAgICAgICBzZXRJbnB1dFZhbHVlKCIuZW52LXdyYXAgLmVudi12YWx1ZSIsICIiKTsKICAgICAgICAgICAgcmV0dXJuIGFsZXJ0KCLliKDpmaTmiJDlip8iKTsKICAgICAgICAgIH0KICAgICAgICAgIGFsZXJ0KGDliKDpmaTlpLHotKU6ICR7cmVzLnN0YXR1c1RleHR9YCk7CiAgICAgICAgfSkKICAgICAgICAuY2F0Y2goKGVycikgPT4gewogICAgICAgICAgYWxlcnQoYOWIoOmZpOWksei0pTogJHtlcnJ9YCk7CiAgICAgICAgfSk7CiAgICB9OwoKICAgIC8vIOiuvue9rum7keWQjeWNlSBpcAogICAgY29uc3Qgc2V0QmxhY2tJcCA9ICgpID0+IHsKICAgICAgY29uc3QgYmxhY2tJcCA9IGdldElucHV0VmFsdWUoIi5ibGFjay1pcC13cmFwIC5ibGFjay1pcCIpOwogICAgICBpZiAoYmxhY2tJcCA9PT0gIiIpIHsKICAgICAgICByZXR1cm4gYWxlcnQoIuWPguaVsOS4jeiDveS4uuepuiIpOwogICAgICB9CiAgICAgIGNvbnN0IHsgc2VjcmV0LCBob3N0IH0gPSBnZXRCYXNlUGFyYW1zKCk7CiAgICAgIGNvbnN0IHVybCA9IGAke2hvc3R9L2JsYWNrX2lwP2lwPSR7YmxhY2tJcH0mc2VjcmV0PSR7c2VjcmV0fWA7CiAgICAgIGZldGNoVXJsKAogICAgICAgIHVybCwKICAgICAgICAoKSA9PiBhbGVydCgi5re75Yqg5oiQ5YqfIiksCiAgICAgICAgKGVycikgPT4gYWxlcnQoYOa3u+WKoOWksei0pTogJHtlcnJ9YCkKICAgICAgKTsKICAgIH07CgogICAgLy8g6K6+572u5Zyw5Z+f55m95ZCN5Y2VCiAgICBjb25zdCBzZXRXaGl0ZUFyZWEgPSAoYWN0aW9uID0gInNldCIpID0+IHsKICAgICAgY29uc3Qgd2hpdGVBcmVhID0gZ2V0SW5wdXRWYWx1ZSgiLndoaXRlLWFyZWEtd3JhcCAud2hpdGUtYXJlYSIpOwogICAgICBpZiAod2hpdGVBcmVhID09PSAiIikgewogICAgICAgIHJldHVybiBhbGVydCgi6L6T5YWl5Y+C5pWw5LiN6IO95Li656m6Iik7CiAgICAgIH0KICAgICAgY29uc3QgYWN0aW9uVGlwID0gYWN0aW9uID09PSAic2V0IiA/ICLmt7vliqAiIDogIuenu+mZpCI7CiAgICAgIGNvbnN0IHsgc2VjcmV0LCBob3N0IH0gPSBnZXRCYXNlUGFyYW1zKCk7CiAgICAgIGNvbnN0IHVybCA9IGAke2hvc3R9L3doaXRlX2FyZWEvJHthY3Rpb259P2FyZWE9JHt3aGl0ZUFyZWF9JnNlY3JldD0ke3NlY3JldH1gOwogICAgICBmZXRjaFVybCgKICAgICAgICB1cmwsCiAgICAgICAgKCkgPT4gYWxlcnQoYCR7YWN0aW9uVGlwfeaIkOWKn2ApLAogICAgICAgIChlcnIpID0+IGFsZXJ0KGAke2FjdGlvblRpcH3lpLHotKU6ICR7ZXJyfWApCiAgICAgICk7CiAgICB9OwogIDwvc2NyaXB0Pgo8L2h0bWw+Cg==`
)

const (
	Gin_IpAddrInfoKey = "ip_addr_info" // 在 gin 上下文中存放 IP 地址信息的键
)
