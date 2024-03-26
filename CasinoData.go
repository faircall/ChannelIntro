package main

func MakeCasinoData() Scene {
	var result Scene
	var tileMaps []TileMap
	tileWidth := 16

	tilesAcross := 960 / tileWidth
	tilesDown := 512 / tileWidth
	tileMap := TileMap{Tiles: nil, Width: tilesAcross, Height: tilesDown, TileWidth: tileWidth, TileHeight: tileWidth}
	for i := 0; i < tilesAcross*tilesDown; i++ {
		tileMap.Tiles = append(tileMap.Tiles, i)
	}

	tileMaps = append(tileMaps, tileMap)
	result.TileMap = tileMaps

	return result

}
