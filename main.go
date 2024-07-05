package main

import (
	"fmt"
	"bufio"
	"os"
	"math/rand"
	"time"
	"net/http" 
	"encoding/json"
	"io/ioutil"
	"github.com/McFuddy2/Pokedex/pokecache"
)


type Config struct {
	BaseURL	 string	`json:base_url`
    Next 	 string `json:"next"`
    Previous string `json:"previous"`
}

var cache *pokecache.Cache

func fetchLocationData(url string) ([]byte, error) {
	if val, ok := cache.Get(url); ok {
		fmt.Println("... ... ...Cache hit!")
		return val, nil
	}

	fmt.Println("... ... Oh! it got away, Cache missed! Fetching from API")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}

	cache.Add(url,body)
	return body, nil
}


func parseLocationData(data []byte) ([]string, *Config, error) {
    var result struct {
        Results []struct {
            Name string `json:"name"`
        } `json:"results"`
        Next     string `json:"next"`
        Previous string `json:"previous"`
    }

    if err := json.Unmarshal(data, &result); err != nil {
        return nil, nil, err
    }

    locations := make([]string, 0, len(result.Results))
    for _, loc := range result.Results {
        locations = append(locations, loc.Name)
    }

    config := &Config{
        Next:     result.Next,
        Previous: result.Previous,
    }

    return locations, config, nil
}


func fetchLocations(config *Config, direction string) ([]string, *Config, error) {
	var url string
	if direction == "next" {
		if config.Next != "" {
			url = config.Next
		} else {
			url = config.BaseURL
		}
	} else if direction == "previous" {
		if config.Previous != ""{
			url = config.Previous
		} else {
			fmt.Println("Silly Swanna, we are allready at the beginning of the list. Ill just give you the first page again.")
			url = config.BaseURL
		}
	} else {
		url = config.BaseURL
	}

	data, err := fetchLocationData(url)
	if err != nil {
		return nil, nil, err
	}

	locations, newConfig, err := parseLocationData(data)
	if err != nil {
		return nil, nil, err
	}

	return locations, newConfig, nil
	}


func displayLocations(locations []string) {
	for _, location := range locations {
		fmt.Println(location)
	}
}



func main() {
	cache = pokecache.NewCache(5*time. Minute)
	scanner := bufio.NewScanner(os.Stdin)
	config := &Config{
		BaseURL: "https://pokeapi.co/api/v2/location",
	} 
	for {
		fmt.Print("pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			fmt.Println("You Entered: ", input)

			switch input {
			case "exit":
				fmt.Println("Goodbye for now! I'll CATCH you later!")
				return
			case "joke":
				jokes := []string{
					"Why did the Pikachu cross the road? To get to the other Psyduck!",
					"What do you call a Pikachu that can fix computers? A Geek-achu!",
					"Why did Meowth go to school? To improve his Meowthematics!","	What do you call a lazy Pikachu? A shockslacker.	",
					"Did you hear about the fire-breathing Charizard who got into modeling? He’s a real hot commodity.",
					"Why did the Ditto go to therapy? It was having an identity crisis.",
					"Why did Squirtle join a band? He wanted to be a shellist.",
					"What does a Pokemon scientist use to see invisible Pokemon? A Psyducktape.",
					"Why did the trainer go on strike? The gym leader wasn’t paying him enough.",
					"What do you call a Pokemon that can’t move? Paralyzachu.",
					"What does a Jigglypuff use to style its hair? A pooffant.",
					"Why did the Voltorb stop rolling? It was tired of being round.",
					"Why did Meowth join a cooking class? It wanted to become the top chef at the Meow Meow restaurant.",
					"What did the Marill say after winning a race? I’m feeling quite blue-tiful.",
					"How do you find a lost Pokemon? You just have to look in the right Eevee spot.",
					"What do you call a Parasect that’s really good at math? A fungi-nerd.",
					"Why did the Haunter join a comedy club? It wanted to be a real ghost comedian.",
					"What did the Machamp say before pumping iron? I’m going to give these weights a real hitmon-lee.",
					"Why did the Slowpoke decide to start working out? It wanted to evolve into a Slowbro-swimmer.",
					"What’s a Frogadier’s favorite type of sandwich? A deluxe fly and croak-a-mole.",
					"What did the Cubone say when it found its missing bone? It was a real skull-streak to find this.",
					"Why did the Bulbasaur cross the road? To get to the other seed-oak tree.",
					"How does a Pokemon apologize for being late? Sorry, I got lost in the Wigglytuff maze.",
					"Why did Pikachu go to the doctor? He was feeling a bit Pikachu.",
					"What did the grass type Pokémon say when he went to the gym? I’m ready to leaf this place feeling stronger!",
					"I used to think I was an electric type, but it turns out I was just shocked all the time.",
					"Why did the water type Pokémon go to the beach? He needed a quick squirtle!",
					"How do you know when a Blastoise is happy? He shell be smiling!",
					"Why did the fire type Pokémon go to the convenience store? He wanted a Charmander bar.",
					"What do you call a Pikachu that can fix anything? A toolchu!",
					"How is a Jigglypuff like a laptop? They both have a power chord.",
					"What do you call a Snorlax with a cold? A Snorelaxative.",
					"Why did Charizard become an actor? He had a flare for the dramatic.",
					"How does a Ditto introduce itself? Ditto, nice to meet you!",
					"Why did the Bulbasaur go to therapy? He needed to work on his bulb-esteem.",
					"How does a Slowpoke tell time? He looks at his Shellder clock.",
					"What did the Squirtle say when he stepped on a Lego? Squirt!",
					"Why did the Abra stop telling jokes? Because he kept making people Drowzee.",
					"How do you get a Pikachu on a bus? You Pokémon!",
					"Why was Gyarados always angry? Because he was schooling all the time.",
					"What do you call a group of Pikachu on vacation? A tropical T-bolt!",
					"Why did Misty get so mad when Psyduck fainted? Because he always let her down.",
					"How does Nidoking clean his castle? He uses a male Nidoqueen!",
					"Q: Why did the Pikachu go on strike? A: It wasn’t getting its fair share of the pikapie!",
					"Q: How does a Jigglypuff keep its hair so perfect? A: With a lot of hair-spray-on!",
					"Q: What do you call a group of Slowpokes? A: A sluggish squad.",
					"Q: What happened when the Magikarp tried to learn how to fly? A: It just kept floundering!",
					"Q: How does a Snorlax stay in good shape? A: It lifts heavy weights.",
					"Q: Why did the Squirtle join a rock band? A: Because it was shell-shocked by the music.",
					"Q: Why was the Charizard kicked out of the library? A: It was constantly booking it!",
					"Q: How does a Gyarados make its bed? A: With sheets of water.",
					"Q: What is a Snorlax’s favorite type of music? A: Lullabys!",
					"Q: How does a Zubat take a photograph? A: With a bat-tery!",
					"Q: What do you call a Snorlax that likes to cook? A: A snoozin’ chef!",
					"Q: Why did the Squirtle cross the road? A: To get to the other shell!",
					"Q: How does a Bulbasaur always win races? A: It has a grass advantage!",
					"Q: Why did the Machop start a gym? A: Because it was a body builder!",
					"Q: What do you call a Psyduck’s psychic powers? A: Brainquack!",
					"Q: Why did the Geodude take a job at a bakery? A: To rock the dough!",
					"Q: What happens when a Pikachu eats too much candy? A: It gets a shocking sugar rush!",
					"Q: Why did the Jigglypuff go on a world tour? A: To sing its praises!",
					"Q: What is a Ditto’s favorite type of book? A: A copycat novel!",
					"Q: How does a Pikachu keep its fur so soft? A: It uses static guard.",
					"Why did the Pikachu go to the doctor? Because he was feeling a little shocked!",
					"How does a Squirtle hide from its enemies? By using its shell-phone!",
					"What do you call a Pokémon that is always getting lost? A Snor-lax!",
					"I was going to tell you a joke about Charizard, but it was too hot to handle.",
					"Why are Electric-type Pokémon always so jumpy? Because they’re full of volts!",
					"How do you make Pikachu laugh? Tell him a Jiggly-pun!",
					"Why did Ash always take Bulbasaur hiking? Because he was a fungi to be around!",
					"What type of music does Jigglypuff listen to? Pop music, of course!",
					"Why did the Chansey go to the doctor? Because it had an egg-cellent sense of humor!",
					"How do you fix a broken Pokémon? With a Pikachu!",
					"What did the trainer say when Pikachu refused to battle? Fine, have it your Whey!",
					"Why is Jynx always wearing lipstick? Because she wants to kiss every Pokémon she meets!",
					"What do you get when you cross a Slowpoke with a Snorlax? A Slow-snor, of course!",
					"What did the Magikarp say when it finally evolved into a Gyarados? I’ve made a big Splash in the Pokémon world!",
					"Why did the Pokémon go to the beach? To catch some Seel-fies!",
					"How do you know if a grass-type Pokémon is happy? It starts to Bulba-smile!",
					"What did the trainer say when Psyduck couldn’t remember any moves in battle? It’s okay, we all have our Pikachu-ps!",
					"Why was Snorlax always late to the Pokémon gym? Because he was lax with his time management!",
					"What did Pikachu say when it got a cold? Pika-achoo!",
					"How do you make a Pokémon trainer angry? Tell them their favorite Pokémon is Gone-saur!",
					"There’s no such thing as too much Pikachu in your life.",
					"Why catch ’em all when you can just buy ’em all on eBay?",
					"My Pokemon gym badge collection consists of mostly gym passes to the snack bar.",
					"I don’t always battle trainers, but when I do, I make sure they’re lower level than me.",
					"You haven’t lived until you’ve heard a grown man yell ‘I choose you!’ in public.",
					"If I had a dollar for every time someone asked if I was playing PokemonGo, I’d have enough money to buy a Master Ball.",
					"I may not have a six-pack, but I do have a level 100 Blastoise.",
					"The real struggle is trying to decide which starter Pokemon to choose.",
					"Sure, I’ve heard of leg day, but have you tried catching wild Pokemon all day?",
					"I can finally cross ‘trying to catch a Pokemon with my toes’ off my bucket list.",
					"I don’t always use potions, but when I do, it’s to heal my Pokemon’s emotional wounds.",
					"If you’re having a bad day, just remember that somewhere out there, a Magikarp is still trying to learn Splash.",
					"I’m pretty sure my Pokemon Go addiction is just a clever disguise for my exercise avoidance.",
					"Forget the Avengers, the true dream team is a Pikachu and a Charizard.",
					"I’m not saying Team Rocket was right, but they did have some fly outfits.",
					"Not all heroes wear capes, some have mastered the art of breeding shiny Pokemon.",
					"My mom told me I can’t bring my Pokemon to college, but she didn’t say anything about my imaginary ones.",
					"I don’t believe in love at first sight, but I do believe in catching a rare shiny Pokemon on the first throw.",
					"Gotta catch ’em all? More like gotta nap ’em all.",
					"I can’t adult today, I have to evolve my Eevee.",
					"A Pikachu in the hand is worth two in the tall grass.",
					"A Ditto in the Pokeball is worth a thousand in the Pokedex.",
					"A Snorlax’s snores are louder than a Rhyhorn’s stampede.",
					"A Pidgeot never goes out of style, but its feathers do.",
					"Beware of Dragonite in sheep’s clothing.",
					"Don’t judge a Magikarp by its splash.",
					"An Eevee evolvement a day keeps the trainers at bay.",
					"It’s better to get a Zapdos than to get zapped, Dos.",
					"A Clefairy in the park is worth two in the dark.",
					"When life gives you Voltorb, just roll with it.",
					"A Slowpoke may be slow, but its wisdom is timeless.",
					"A Jigglypuff’s lullaby can put even a Snorlax to sleep.",
					"It’s always sunny in Fire-type gyms.",
					"A Bulbasaur a day keeps Team Rocket away.",
					"Carbink may be rare, but its cuteness is abundant.",
					"A Chansey on your team brings good luck, and extra HP.",
					"Don’t let a Grimer get a grip on you, it’s a real sticky situation.",
					"The early trainer catches the legendary bird.",
					"The key to catching a Mew is to stay Pika-quiet.",
					"A Gyarados may be intimidating, but it’s just a Magikarp with attitude.",
					"I might be a Jynx, but I’ll still make your Psyduck.",
					"I’ll let you Pikachu my heart any day.",
					"Do you want to see my Magikarp trick?",
					"I’m always down for a good Squirtle fight.",
					"No need to be afraid of my Gyarados, it only bites between the sheets.",
					"Is that a Voltorb in your pocket or are you just happy to see me?",
					"I’m feeling a bit Machop today, are you up for a challenge?",
					"Charmander, charmander, let’s make things hot and steamy.",
					"I may be a Snorlax, but I’ll still give you the ride of your life.",
					"I can’t resist your Charmeleon.",
					"I hope you don’t mind a little Geodude pun.",
					"My Metapod is ready to Harden for you.",
					"You must be a Water-type because I’m feeling a bit Squirtle-y.",
					"I may be a trainer, but you can be my master anytime.",
					"I’ll show you my Pikachu if you show me your Raichu.",
					"I’m not just a trainer, I’m also a Charmer.",
					"You must have a Grass-type Pokemon, because you’re making me Bulbasaur.",
					"I may be an Eevee, but I have a lot of potential to evolve.",
					"I may be an Eevee, but I have a lot of potential to evolve.",
					"It’s always sunny in Fire-type gyms.”",
					"Why did the Voltorb stay hidden in the grass? Because it was feeling a bit Shellder-shy.",
					"Did you hear about the Pikachu that started his own business? It was an electric venture.",
					"What do you call a Pokemon who loves to bake? A Fluffy Metapod.",
					"How does a Geodude keep his hair in place? With Rock-hard style gel.",
					"Why did the Snorlax go on a diet? So he could fit into his Swellow outfit.",
					"What do you get when you cross a Bulbasaur with a Chansey? A Bulba-baby.",
					"Did you hear about the Eevee who took up gardening? It evolved into a Leafron.",
					"How do you explain a Slowpoke’s laid-back attitude? He’s just trying to take things one Shellder at a time.",
					"What kind of music do Pokemon listen to? Anything with a good Charizard-one it.",
					"What do you call a Snorlax who loves to dance? A Disco-nap.",
					"How do you catch a flying Jigglypuff? With a Puff-ball.",
					"Did you hear about the Slowbro who opened a daycare? He’s in charge of Slow-growth development.",
					"What’s a Ditto’s favorite hobby? Reflecting on life.",
					"Why was the Snorlax feeling down? He forgot to set his Alarming clock.",
					"What kind of car does Pikachu drive? A Voltswagen.",
					"Why did the Charmander go to the doctor? He was feeling a bit Char-zzy.",
					"How does a Trainer know when their Pokedex is overheating? It starts to Squirtle steam.",
					"What do you call a dance party with Pokemon? A Pika-party.",
					"Why was the Clefairy feeling guilty? It accidentally used Metronome in a Silentforest.",
					"What did the Squirtle say to the grass-type Pokemon? Lettuce have a battle!",
					"Knock, knock. Who’s there? Pikachu. Pikachu who? Pikachu in the eye, it’ll make you cry!",
					"Knock, knock. Who’s there? Jigglypuff. Jigglypuff who? Jigglypuff up your pillow before you go to bed!",
					"Knock, knock. Who’s there? Charizard. Charizard who? Charizard-nado coming your way!",
					"Knock, knock. Who’s there? Bulbasaur. Bulbasaur who? Bulbasaur makes your garden grow!",
					"Knock, knock. Who’s there? Snorlax. Snorlax who? Snorlax and relax, it’s the weekend!",
					"Knock, knock. Who’s there? Squirtle. Squirtle who? Squirtle when I tickle you!",
					"Knock, knock. Who’s there? Psyduck. Psyduck who? Psyduck you so hard, you won’t even know what hit you!",
					"Knock, knock. Who’s there? Eevee. Eevee who? Eevee’s dropping from the sky!",
					"Knock, knock. Who’s there? Meowth. Meowth who? Meowth-tiful day for a picnic, don’t you think?",
					"Knock, knock. Who’s there? Togepi. Togepi who? Togepi-rama, it’s party time!",
					"Knock, knock. Who’s there? Gengar. Gengar who? Gengar-ling with my best friends!",
					"Knock, knock. Who’s there? Mewtwo. Mewtwo who? Mewtwo hilarious to handle!",
					"Knock, knock. Who’s there? Machop. Machop who? Machop chop chop, let’s battle!",
					"Knock, knock. Who’s there? Ditto. Ditto who? Ditto-tally awesome, that’s what I am!",
					"Knock, knock. Who’s there? Vulpix. Vulpix who? Vulpix in a blanket, please!",
					"Knock, knock. Who’s there? Jynx. Jynx who? Jynx and you’ll never find me!",
					"Knock, knock. Who’s there? Pidgey. Pidgey who? Pidgey-back ride, anyone?",
					"Knock, knock. Who’s there? Onix. Onix who? Onix-pected visitor!",
					"Knock, knock. Who’s there? Growlithe. Growlithe who? Growlithe to help me catch this Caterpie!",
					"Knock, knock. Who’s there? Geodude. Geodude who? Geodude, I love you!",
					
				}
				rand.Seed(time.Now().UnixNano())
				randomIndex := rand.Intn(len(jokes))
				fmt.Println(jokes[randomIndex])

			case "help":
				fmt.Println("Hello, and welcome to your Pokedex! My name is Puffluff your AI assistant! but you can call me Puff!")
				fmt.Println("Looks like you are looking for some help! Let me find a list of some commands you can try!")
				fmt.Println("")
				fmt.Println("help: Displays a 'help' message")
				fmt.Println("exit: Exits the Pokedex")
				fmt.Println("map: this will display the next 20 locations for you to be able to choose from")
				fmt.Println("mapb: this will display the PREVIOUS 20 locations")
				fmt.Println("joke: I will tell you a pokemon joke that's sure to leave you Weezing")
				
			case "map":
				locations, newConfig, err := fetchLocations(config, "next")
				if err != nil {
					fmt.Println("Error fetching locations:", err)
				} else {
					displayLocations(locations)
					config = newConfig
				}
			case "mapb":
				locations, newConfig, err := fetchLocations(config, "previous")
				if err != nil {
					fmt.Println("Error fetching locations:", err)
				} else {
					displayLocations(locations)
					config = newConfig
				}
			default:
				fmt.Println("Unknown command:", input)
			}
		}
	}
}