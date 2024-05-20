# mboff

Mboff is tiny CLI that allows the optimisation of MBTiles by removing unnecessary data. You only need to provide it with a key/value pair and `mboff` will search and remove such data.

## Installation:

```bash
brew tap jjcfrancisco/mboff # Adds the Github repository as a tap
brew install yeo
```

## Usage
`mboff` requires you to at least give a path to the existing *MBTiles* file and the key/value pair of the data you wish to remove from the file. Optionally, you may want to filter the removal of data by zoom level:
```bash
mboff [file path] [key value pair] [zoom level]
```
Examples:
```bash
# Remove data that contains category=road key/value pair
mboff myMap.mbtiles category=road

# Remove data that contains category=road key/value pair in zoom level 10
mboff myMap.mbtiles category=road 10
```

## Future
* A command for *checking* if a key/value pair exist in a MBTiles file and how many times is present.
* Extend/add validation for args

## License

See [`LICENSE`](./LICENSE)
