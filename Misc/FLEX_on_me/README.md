[< Go to Misc](/Misc) OR [<< Go to Home](/)
# FLEX on me - 150 Points
## Problem
I know you Gen-Z kids don't know what a [pager](https://en.wikipedia.org/wiki/Pager) is, but I was browsing the 900MHz band searching for a friend. I came across [this message](FLEX_on_me.wav). Can you decode the POCSAG FLEX message? Tell me the name of the doctor no longer covering in format last, first.

Example: `Truluck, Charles`

## Hint
I'm only giving you Gen-Z-ers a hint because you're new to the paging game. \
If you are a Windows gamer, try PDW. \
If you prefer the professional Mac architecture, try Multimon-ng.

## Solution
I made sure this problem worked both on Windows and Mac before deploying it.

#### MacOS
Since I primarily use a Mac, I took the following Multimon-ng command: `multimon-ng -t wav -a FLEX flex_on_me.wav` which outputted the decoded messages below.

#### Windows
I tested PDW on my ***gamer rig***, which worked equally as flawlessly as Multimon-ng. The only unintended side-effect was needing to install a third-party internal audio router (I chose [VB-Audio](https://www.vb-audio.com/Cable/)), but there may be an easy way to work around this with PDW. I fed the audio into VB-Audio, and set PDW to listen to that input as the discriminator.

```
FLEX: 2019-11-19 19:08:11 1600/2/K/A 07.097 [000002607] ALN TEST
FLEX: 2019-11-19 19:08:12 1600/2/K/A 07.098 [000006324] ALN N5@Hc4jE682ZPH;^COO7M2
FLEX: 2019-11-19 19:08:12 1600/2/K/A 07.099 [000006773] ALN F<>EBM`YQOZ[5IA6@;>?_bo87;iHDOf>?=AJ7aU8E\S0hl:Edi>QRVdCdZ2=X4VO0QdH5CZj=Wc`V0JA`WC21X^TAF4VCJ\oeQhPk2Uh6BCjNfi`15kXnFUB073ShUj]7kTFeTSHoR;]@K4UX]B5de1G?5M6obi>8oRMPo?HMH;
FLEX: 2019-11-19 19:08:12 1600/2/K/A 07.100 [000001209] ALN ID 1-1988 [Kish, Lauren] is no longer covering for ID 17201 [PROSTHETIC & ORTHOTIC SERVICE]. Please do not reply to this message.
```

## Flag
`Kish, Lauren`
