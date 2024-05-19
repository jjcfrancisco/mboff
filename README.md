# mboff

Mboff is tiny CLI that allows the optimisation of MBTiles by removing unnecessary data. You only need to provide it with a key/value pair and `mboff` will search and remove such data.

## Usage
`mboff` requires the user to give a path to the existing *MBTiles* file and the key/value pair of the data you wish to remove from the file:
```bash
mboff [file path] [key value pair]
```
Example:
```bash
mboff myMap.mbtiles category=road
```

## Future
* Add mboff to Homebrew
* A command for *checking* if a key/value pair exist in a MBTiles file and how many times is present.
* Extend/add validation for args

## License

See [`LICENSE`](./LICENSE)
