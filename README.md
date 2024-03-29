# Metronome


https://user-images.githubusercontent.com/16705876/228581110-e1ca806a-86d3-4ae7-83a6-c364d8bc8446.mov



(Put the sound on 😊)

## Background
* I wanted to build out a **metronome** on the command line so I didn't have to rely on Google's metronome
* I used the Go [beep](https://github.com/faiface/beep) package to output sound
* I created a [Homebrew tap](https://github.com/MadeleineSmith/homebrew-metronome) for [easy installation](#first-install-instructions) of the metronome package

---

## First install instructions:
* ` brew tap madeleinesmith/metronome && brew install metronome `

---

## How to release new version and update local package:
Releasing new version:
* Tag code using ` git tag -a v1.0.0 -m "version 1.0.0" `
* Push tag with ` git push origin v0.2.0 `
* Create a new release on [GitHub](https://github.com/MadeleineSmith/metronome/releases/new) for that tag
* Copy the link of the `tar.gz` file on GitHub
* Change the `url` line of `homebrew-metronome` [repo](https://github.com/MadeleineSmith/homebrew-metronome/blob/4661e8c8d8ef9dcafb2a46e645d57550990ba31b/metronome.rb#L7) to be this
* And also update the `sha256` line by running `shasum -a 256 xxxxxxx.tar.gz` on the downloaded tar file (above)  
* Commit and push the `homebrew-metronome` repo with these edits

Updating the local package:
* Run ``` brew update && brew upgrade metronome ```

---

## Usage instructions
When using the Brew package run:
* `metronome -beats-per-minute=a -beats-per-bar=b -subdivisions=c,d,e,f`

e.g.
* `metronome -beats-per-minute=15 -beats-per-bar=4 -subdivisions=4,4,4,7`


If running locally use:
* `go run main.go -beats-per-minute=a -beats-per-bar=b -subdivisions=c,d,e,f`

---

## Tutorials referenced whilst building:
Creating a Homebrew tap:
* https://betterprogramming.pub/a-step-by-step-guide-to-create-homebrew-taps-from-github-repos-f33d3755ba74
* https://flowerinthenight.com/blog/2019/07/30/homebrew-golang

Using go:embed:
* https://blog.jetbrains.com/go/2021/06/09/how-to-use-go-embed-in-go-1-16/
