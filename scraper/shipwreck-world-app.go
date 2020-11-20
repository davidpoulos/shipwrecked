package scraper

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/davidpoulos/shipwrecked/shipwreck"
	"github.com/gocolly/colly/v2"
)

// Website
var ShipWreckCrawlSites = []string{
	"https://www.shipwreckworld.com/maps/17-fathom-wreck",
	"https://www.shipwreckworld.com/maps/merida",
	"https://www.shipwreckworld.com/maps/a-e-vickery-schooner-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/mfv-zeila-l758",
	"https://www.shipwreckworld.com/maps/a-r-noyes-canal-boat-lake-champlain-shipwreck",
	"https://www.shipwreckworld.com/maps/milan-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/acme-tug-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/munson-dredge-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/acme-propeller-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/muscallonge-tug-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/admiral",
	"https://www.shipwreckworld.com/maps/mv-vashon",
	"https://www.shipwreckworld.com/maps/air-boat-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/mv-wilhem-gustloff",
	"https://www.shipwreckworld.com/maps/alabama-steamer-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/niagara-1",
	"https://www.shipwreckworld.com/maps/alberta-tug-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/nimrod",
	"https://www.shipwreckworld.com/maps/alexandra-steamer-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/nineveh",
	"https://www.shipwreckworld.com/maps/algeria",
	"https://www.shipwreckworld.com/maps/north-carolina",
	"https://www.shipwreckworld.com/maps/aloha-schooner-barge-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/north-colborne-island-barge-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/america-barge-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/o-j-walker-schooner-lake-champlain-shipwreck",
	"https://www.shipwreckworld.com/maps/andrew-b",
	"https://www.shipwreckworld.com/maps/o-w-cheny",
	"https://www.shipwreckworld.com/maps/annabell-wilson-schooner-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/ocean-wave-paddlewheel-steamer-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/annie-falconer-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/oconto-steamer-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/apps-2-barge-oneida-lake",
	"https://www.shipwreckworld.com/maps/olive-branch-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/arizona-steamer-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/oliver-mowat-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/atlantic",
	"https://www.shipwreckworld.com/maps/oneida",
	"https://www.shipwreckworld.com/maps/atlasco-steamer-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/oneida-lake-sunken-barge",
	"https://www.shipwreckworld.com/maps/aycliffe-hall",
	"https://www.shipwreckworld.com/maps/ontario",
	"https://www.shipwreckworld.com/maps/banshee-propeller-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/orcadian-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/barge",
	"https://www.shipwreckworld.com/maps/oxford",
	"https://www.shipwreckworld.com/maps/barge-f",
	"https://www.shipwreckworld.com/maps/pascal-p-pratt",
	"https://www.shipwreckworld.com/maps/barge-43-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/passaic-propeller-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/barge-f-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/persian",
	"https://www.shipwreckworld.com/maps/belle-sheridan-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/phalarope",
	"https://www.shipwreckworld.com/maps/betty-hedger",
	"https://www.shipwreckworld.com/maps/philip-d-armour",
	"https://www.shipwreckworld.com/maps/bishops-derirck-barge-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/phoenix-steamer-lake-champlain-shipwreck",
	"https://www.shipwreckworld.com/maps/briton",
	"https://www.shipwreckworld.com/maps/prince-regent-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/brown-brothers",
	"https://www.shipwreckworld.com/maps/ps-charles-h-spencer",
	"https://www.shipwreckworld.com/maps/brunswick-propeller-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/ps-olympian",
	"https://www.shipwreckworld.com/maps/c-b-benson-1",
	"https://www.shipwreckworld.com/maps/psyche-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/c-w-elphicke",
	"https://www.shipwreckworld.com/maps/queen-of-the-west",
	"https://www.shipwreckworld.com/maps/canobie",
	"https://www.shipwreckworld.com/maps/quinte-steamer-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/carlingford-schooner-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/r-h-rae-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/catherine-yacht-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/raleigh",
	"https://www.shipwreckworld.com/maps/cecil-j",
	"https://www.shipwreckworld.com/maps/raymond-yacht-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/champlain-steamer-lake-champlain-shipwreck",
	"https://www.shipwreckworld.com/maps/rms-atlantic",
	"https://www.shipwreckworld.com/maps/charles-b-hill",
	"https://www.shipwreckworld.com/maps/rms-carpathia",
	"https://www.shipwreckworld.com/maps/charles-foster",
	"https://www.shipwreckworld.com/maps/rms-empress-of-ireland",
	"https://www.shipwreckworld.com/maps/charles-h-davis",
	"https://www.shipwreckworld.com/maps/rms-laconia",
	"https://www.shipwreckworld.com/maps/china-steamer-false-duck-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/rms-lusitania",
	"https://www.shipwreckworld.com/maps/chippewa-propeller-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/rms-rhone",
	"https://www.shipwreckworld.com/maps/city-of-rome",
	"https://www.shipwreckworld.com/maps/rms-titanic",
	"https://www.shipwreckworld.com/maps/city-of-sheboygan-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/robert",
	"https://www.shipwreckworld.com/maps/clara-white-schooner-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/robert-gaskin-barque-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/cleveco",
	"https://www.shipwreckworld.com/maps/rothesay-paddlewheeler-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/cleveland-barge-oneida-lake-wreck",
	"https://www.shipwreckworld.com/maps/roy-a-jodrey-self-loader-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/colonial",
	"https://www.shipwreckworld.com/maps/s-m-douglas-steamer-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/comet-steamer-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/s-k-martin",
	"https://www.shipwreckworld.com/maps/conestoga-freighter-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/sam-cooke-schooner-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/cormorant-tugboat-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/sand-barge-oneida-lake-wreck",
	"https://www.shipwreckworld.com/maps/cornwall-steamer-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/sand-merchant",
	"https://www.shipwreckworld.com/maps/cracker-wreck",
	"https://www.shipwreckworld.com/maps/schooner-three-masted-unknown",
	"https://www.shipwreckworld.com/maps/craftsman",
	"https://www.shipwreckworld.com/maps/schooner-unknown",
	"https://www.shipwreckworld.com/maps/crystal-wreck",
	"https://www.shipwreckworld.com/maps/schooner-unknown-erieau",
	"https://www.shipwreckworld.com/maps/dacotah-propeller-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/schooner-c-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/dauntless-yacht-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/schooner-g-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/david-w-mills-steamer-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/schooner-unidentified-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/dawn",
	"https://www.shipwreckworld.com/maps/scourge-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/dean-richmond",
	"https://www.shipwreckworld.com/maps/shackelton-barge-oneida-lake-wreck",
	"https://www.shipwreckworld.com/maps/diesel-barge-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/siberia",
	"https://www.shipwreckworld.com/maps/drill-rig",
	"https://www.shipwreckworld.com/maps/silgo-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/duke-luedtke",
	"https://www.shipwreckworld.com/maps/sir-c-t-van-straubenzie",
	"https://www.shipwreckworld.com/maps/dundee",
	"https://www.shipwreckworld.com/maps/sir-robert-peel-steamer-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/dupuis-10",
	"https://www.shipwreckworld.com/maps/sloop-island-canal-boat-lake-champlain-shipwreck",
	"https://www.shipwreckworld.com/maps/eastcliffe-hall-freighter-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/smith",
	"https://www.shipwreckworld.com/maps/echo-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-america",
	"https://www.shipwreckworld.com/maps/eduard-bohlen",
	"https://www.shipwreckworld.com/maps/ss-andrea-doria",
	"https://www.shipwreckworld.com/maps/effie-may-trawler-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-arctic",
	"https://www.shipwreckworld.com/maps/egan-sisters-newspaper-barge-oneida-lake",
	"https://www.shipwreckworld.com/maps/ss-atlantus",
	"https://www.shipwreckworld.com/maps/elk-schooner-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-bluefields",
	"https://www.shipwreckworld.com/maps/etta-belle-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-cape-fear",
	"https://www.shipwreckworld.com/maps/f-a-meyer",
	"https://www.shipwreckworld.com/maps/ss-chirripo",
	"https://www.shipwreckworld.com/maps/fabiola-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-city-of-chester",
	"https://www.shipwreckworld.com/maps/fannie-l-jones",
	"https://www.shipwreckworld.com/maps/ss-city-of-new-york",
	"https://www.shipwreckworld.com/maps/finch-barge-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-city-of-rio-de-janeiro",
	"https://www.shipwreckworld.com/maps/fleur-marie-brigantine-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-clallam",
	"https://www.shipwreckworld.com/maps/florence-tug-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-columbia",
	"https://www.shipwreckworld.com/maps/forward-motor-launch-lake-george-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-dix",
	"https://www.shipwreckworld.com/maps/frank-e-vigor",
	"https://www.shipwreckworld.com/maps/ss-edmund-fitzgerald",
	"https://www.shipwreckworld.com/maps/fred-mercur-propeller-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-harvard",
	"https://www.shipwreckworld.com/maps/frontenac",
	"https://www.shipwreckworld.com/maps/ss-iberia",
	"https://www.shipwreckworld.com/maps/frontenac-tug-boat-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/andrea-doria",
	"https://www.shipwreckworld.com/maps/general-butler-schooner-lake-champlain-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-leopoldville",
	"https://www.shipwreckworld.com/maps/george-a-marsh-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-la-bourgogne",
	"https://www.shipwreckworld.com/maps/george-c-finney-1",
	"https://www.shipwreckworld.com/maps/ss-monte-carlo",
	"https://www.shipwreckworld.com/maps/george-george-t-davie-steam-barge-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-ohioan",
	"https://www.shipwreckworld.com/maps/george-whelan",
	"https://www.shipwreckworld.com/maps/ss-oregon",
	"https://www.shipwreckworld.com/maps/grand-view-steamer-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-pacific",
	"https://www.shipwreckworld.com/maps/grenadier-island-sailing-ships-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-palo-alto",
	"https://www.shipwreckworld.com/maps/grenadier-island-wreck-1-and-2-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-pendleton",
	"https://www.shipwreckworld.com/maps/grindstone-island-cars",
	"https://www.shipwreckworld.com/maps/ss-polias",
	"https://www.shipwreckworld.com/maps/h-a-barr",
	"https://www.shipwreckworld.com/maps/ss-richard-montgomery",
	"https://www.shipwreckworld.com/maps/h-g-cleveland",
	"https://www.shipwreckworld.com/maps/ss-robert-e-lee",
	"https://www.shipwreckworld.com/maps/hamilton-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-san-juan",
	"https://www.shipwreckworld.com/maps/harvey-j-kendall-freighter-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-santa-rosa",
	"https://www.shipwreckworld.com/maps/henry-c-daryaw-freighter-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-sapona",
	"https://www.shipwreckworld.com/maps/henry-clay",
	"https://www.shipwreckworld.com/maps/ss-selma",
	"https://www.shipwreckworld.com/maps/schooner-henry-roney",
	"https://www.shipwreckworld.com/maps/ss-valencia",
	"https://www.shipwreckworld.com/maps/hilda-barge-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/ss-yarmouth-castle",
	"https://www.shipwreckworld.com/maps/hmhs-britannic-his-majestys-hospital-ship-no-g618",
	"https://www.shipwreckworld.com/maps/ss-america-1",
	"https://www.shipwreckworld.com/maps/hmhs-glenart-castle",
	"https://www.shipwreckworld.com/maps/st-louis-schooner-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/hmhs-rewa",
	"https://www.shipwreckworld.com/maps/st-james",
	"https://www.shipwreckworld.com/maps/holiday-point-wreck-sailing-ship-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/st-peter-map",
	"https://www.shipwreckworld.com/maps/homer-warren-map",
	"https://www.shipwreckworld.com/maps/steam-launch-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/horse-powered-ferry-boat-lake-champlain-shipwreck",
	"https://www.shipwreckworld.com/maps/steamer-unknown-near-erieau",
	"https://www.shipwreckworld.com/maps/ida-walker-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/stern-castle",
	"https://www.shipwreckworld.com/maps/indiana",
	"https://www.shipwreckworld.com/maps/steven-f-gale",
	"https://www.shipwreckworld.com/maps/iroquois-hms-anson-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/stone-canal-boat-lake-champlain-shipwreck",
	"https://www.shipwreckworld.com/maps/islander-steamer-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/swallow",
	"https://www.shipwreckworld.com/maps/j-g-mcgrath-schooner-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/t-8-schooner",
	"https://www.shipwreckworld.com/maps/james-a-shrigigley-steamer-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/terrys-tug-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/james-b-colgate",
	"https://www.shipwreckworld.com/maps/thomas-h-barge-1-oneida-lake",
	"https://www.shipwreckworld.com/maps/james-h-reed",
	"https://www.shipwreckworld.com/maps/thomas-h-barge-2-oneida-lake",
	"https://www.shipwreckworld.com/maps/japanese-battleship-yamato",
	"https://www.shipwreckworld.com/maps/tiller-wreck-",
	"https://www.shipwreckworld.com/maps/john-a-macdonald-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/tonawanda-propeller-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/john-b-king-scow-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/trade-wind",
	"https://www.shipwreckworld.com/maps/john-b-cowle",
	"https://www.shipwreckworld.com/maps/tug-unknown-117th-street",
	"https://www.shipwreckworld.com/maps/john-b-lyon",
	"https://www.shipwreckworld.com/maps/tug-elsie-k-wreck-debris-oneida-lake",
	"https://www.shipwreckworld.com/maps/john-b-griffin",
	"https://www.shipwreckworld.com/maps/tug-thomas-h",
	"https://www.shipwreckworld.com/maps/john-j-boland",
	"https://www.shipwreckworld.com/maps/tv-frank-h-buck",
	"https://www.shipwreckworld.com/maps/john-pridgeon",
	"https://www.shipwreckworld.com/maps/tv-lyman-stewart",
	"https://www.shipwreckworld.com/maps/julia-schooner-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/twa-flight-800",
	"https://www.shipwreckworld.com/maps/julia-b-merrill-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/two-fannies",
	"https://www.shipwreckworld.com/maps/juno-steamer-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/u-166",
	"https://www.shipwreckworld.com/maps/k-h-p-barge-wreck-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/u-576",
	"https://www.shipwreckworld.com/maps/katie-eccles-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/undine-schooner-shipwreck-lake-ontario",
	"https://www.shipwreckworld.com/maps/keystorm-freighter-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/us-104-barge-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/king-cruiser-wreck",
	"https://www.shipwreckworld.com/maps/us-coast-guard-boat-56022",
	"https://www.shipwreckworld.com/maps/kinghorn-schooner-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/uss-arizona-bb-39",
	"https://www.shipwreckworld.com/maps/land-tortoise",
	"https://www.shipwreckworld.com/maps/uss-benevolence-ah-13",
	"https://www.shipwreckworld.com/maps/laura-grace-tugboat-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/uss-corry-dd-334",
	"https://www.shipwreckworld.com/maps/lewiston-steamer-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/uss-hornet-cv-8",
	"https://www.shipwreckworld.com/maps/lillie-parsons-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/uss-macon-zrs-5",
	"https://www.shipwreckworld.com/maps/little-wissahickon",
	"https://www.shipwreckworld.com/maps/uss-monitor",
	"https://www.shipwreckworld.com/maps/loblaws-wreck",
	"https://www.shipwreckworld.com/maps/uss-san-diego-ca-6",
	"https://www.shipwreckworld.com/maps/londonderry-dredge-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/uss-susan-b-anthony-ap-72",
	"https://www.shipwreckworld.com/maps/louis-oneil",
	"https://www.shipwreckworld.com/maps/valentine",
	"https://www.shipwreckworld.com/maps/lycoming",
	"https://www.shipwreckworld.com/maps/w-c-richardson-freighter-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/lyman-m-davis-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/war-1812-wreck-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/mabel-wilson",
	"https://www.shipwreckworld.com/maps/washington-irving-schooner-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/maggie-l-schooner-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/water-witch-schooner-lake-shamplain-shipwreck",
	"https://www.shipwreckworld.com/maps/manola-steamer-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/whg-speed-boat-st-lawrence-river-shipwreck",
	"https://www.shipwreckworld.com/maps/manzanilla-schooner-lake-erie-shipwreck",
	"https://www.shipwreckworld.com/maps/william-jamieson-schooner-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/marengo",
	"https://www.shipwreckworld.com/maps/william-stevens",
	"https://www.shipwreckworld.com/maps/mary-kay-tugboat-lake-ontario-shipwreck",
	"https://www.shipwreckworld.com/maps/wilma",
	"https://www.shipwreckworld.com/maps/mecosta",
	"https://www.shipwreckworld.com/maps/wolfe-islander-propeller-lake-ontario-shipwreck",
}

func ScrapeShipWreckWorldSite() {

	var count int
	f, _ := os.Create("shipwrecks.txt")

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		//colly.Debugger(&debug.LogDebugger{}),
	)

	// Find and visit all links
	c.OnHTML("html", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))

		r := ExtractShipWreckInfo(e)

		b, _ := json.Marshal(r)

		f.Write(b)
		f.WriteString("\n")

		count = count + 1

		if count < len(ShipWreckCrawlSites) {
			c.Visit(ShipWreckCrawlSites[count])
		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Limit(&colly.LimitRule{
		RandomDelay: 5 * time.Second,
	})

	fmt.Printf("OG COUNT %d", count)
	c.Visit(ShipWreckCrawlSites[count])
}

func ExtractShipWreckInfo(e *colly.HTMLElement) *shipwreck.Shipwreck {
	// Scrape the fields from the HTML
	wreckName := e.ChildText("#ContentPlaceHolder1_TitleSecondLabel")
	latitude := e.ChildText("#map-description > div > div:nth-child(4) > strong:nth-child(2)")
	longitude := e.ChildText("#map-description > div > div:nth-child(4) > strong:nth-child(3)")
	yearBuilt := e.ChildText("#ContentPlaceHolder1_articleWrapperDiv > div.row.map-times-wrapper > div:nth-child(1) > h4 > span")
	yearSank := e.ChildText("#ContentPlaceHolder1_articleWrapperDiv > div.row.map-times-wrapper > div:nth-child(2) > h4 > span")
	difficultyLevel := e.ChildText("#ContentPlaceHolder1_articleWrapperDiv > div.row.map-times-wrapper > div:nth-child(4) > h4 > span")
	depth := e.ChildText("#ContentPlaceHolder1_articleWrapperDiv > div.row.map-times-wrapper > div:nth-child(3) > h4")

	// second
	if latitude == "" || longitude == "" {
		fmt.Printf("[Coordinates Loc. #2] ")
		latitude = e.ChildText("#map-description > div > div:nth-child(3) > strong:nth-child(2)")
		longitude = e.ChildText("#map-description > div > div:nth-child(3) > strong:nth-child(3)")
	}

	// third
	if latitude == "" || longitude == "" {
		fmt.Printf("[Coordinates Loc. #3] ")
		latitude = e.ChildText("#map-description > div > div:nth-child(5) > strong:nth-child(2)")
		longitude = e.ChildText("#map-description > div > div:nth-child(5) > strong:nth-child(3)")
	}

	return &shipwreck.Shipwreck{
		Name:            CleanName(wreckName),
		Latitude:        CleanLatitude(latitude),
		Longitude:       CleanLongitude(longitude),
		YearSank:        CleanYearSank(yearSank),
		YearBuilt:       CleanYearBuilt(yearBuilt),
		DifficultyLevel: CleanDifficultyLevel(difficultyLevel),
		Depth:           CleanDepth(depth),
		Coordinates:     CleanLatitude(latitude) + "," + CleanLongitude(longitude),
	}
}

func CleanDepth(depth string) string {
	depthContent := []rune(depth)
	return strings.TrimSpace(string(depthContent[5:]))
}

func CleanName(name string) string {
	return name
}

func CleanYearBuilt(year string) int {
	y, _ := strconv.Atoi(year)
	return y
}

func CleanYearSank(year string) int {
	y, _ := strconv.Atoi(year)
	return y
}

func CleanLongitude(long string) string {
	//fmt.Printf("GOT A LONG %s\n", long)
	// START_INFO |    Longitude:   44° 16' 48.8388"

	long = strings.ReplaceAll(long, "Longitude:", "")
	long = strings.TrimSpace(long)

	// Remove ° ' "
	long = strings.ReplaceAll(long, "°", "")
	long = strings.ReplaceAll(long, "'", "")
	long = strings.ReplaceAll(long, "\"", "")

	//[44 16 48.8388 N]
	chunks := strings.Split(long, " ")
	if len(chunks) > 3 {
		long = gpsToDegrees(chunks)
	} else {
		long = ""
	}
	return long
}

func CleanLatitude(lat string) string {
	//	fmt.Printf("GOT A LAT %s\n", lat)
	lat = strings.ReplaceAll(lat, "Latitude:", "")
	lat = strings.TrimSpace(lat)

	// Remove ° ' "
	lat = strings.ReplaceAll(lat, "°", "")
	lat = strings.ReplaceAll(lat, "'", "")
	lat = strings.ReplaceAll(lat, "\"", "")

	// [44 16 48.8388 N]
	chunks := strings.Split(lat, " ")

	if len(chunks) > 3 {
		lat = gpsToDegrees(chunks)
	} else {
		lat = ""
	}

	return lat
}

func CleanDifficultyLevel(dLevel string) string {
	return dLevel
}

// pre: [44 16 48.8388 N]
// post: 44 + 0.26666666666666666 + 0.013566333333333333
//       ~44.27
// Lazy Conversion
func gpsToDegrees(gps []string) string {
	degree, _ := strconv.ParseFloat(gps[0], 64)
	mins, _ := strconv.ParseFloat(gps[1], 64)
	secs, _ := strconv.ParseFloat(gps[2], 64)
	return fmt.Sprintf("%v", degree+(mins/60)+(secs/3600))
}
