package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/johanavril/server-communication-benchmark/pb"
	"google.golang.org/grpc/metadata"
	"net/http"
	"testing"
)

func Benchmark_Ping(b *testing.B) {
	s, err := New()
	if err != nil {
		b.FailNow()
	}
	defer s.TearDown()
	ctx := context.Background()

	b.Run("REST", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s.rest.Get(restAddr + "/ping")
		}
	})

	b.Run("GRPC", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s.grpc.Ping(ctx, &pb.PingRequest{Ping: "PING"})
		}
	})
}

func Benchmark_Small(b *testing.B) {
	s, err := New()
	if err != nil {
		b.FailNow()
	}
	defer s.TearDown()
	ctx := context.Background()

	b.Run("REST", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req := smallRequest{
				Identity: identity{
					Username: "John Doe",
					Email: "john@doe.com",
					Country: "Indonesia",
				},
			}
			jsonBytes, err := json.Marshal(req)
			if err != nil {
				b.FailNow()
			}

			httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, restAddr + "/small", bytes.NewBuffer(jsonBytes))
			if err != nil {
				b.FailNow()
			}

			httpResp, err := s.rest.Do(httpReq)
			if err != nil {
				b.FailNow()
			}

			var resp smallResponse
			if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
				httpResp.Body.Close()
				b.FailNow()
			}

			_ = resp.Summary
		}
	})

	b.Run("GRPC", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req := pb.SmallRequest{
				Identity: &pb.Identity{
					Username: "John Doe",
					Email: "john@doe.com",
					Country: "Indonesia",
				},
			}
			resp, err := s.grpc.Small(ctx, &req)
			if err != nil {
				b.FailNow()
			}

			_ = resp.GetSummary()
		}
	})
}

func Benchmark_Big(b *testing.B) {
	s, err := New()
	if err != nil {
		b.FailNow()
	}
	defer s.TearDown()
	ctx := context.Background()

	b.Run("REST", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req := bigRequest{
				Identity: identity{
					Username: "John Doe",
					Email: "john@doe.com",
					Country: "Indonesia",
				},
				Location: location{
					Lat: 40.8223286,
					Lon: -96.7982002,
				},
				Interest: []string{"fantasy", "sport", "tech", "lifestyle"},
				Bookmark: []*content{
					{
						ID: "111111111111",
						Title: "Then Came the Night",
						Description: "Then came the night of the first falling star. It was seen early in the morning, rushing over Winchester eastward, a line of flame high in the atmosphere. Hundreds must have seen it and taken it for an ordinary falling star. It seemed that it fell to earth about one hundred miles east of him.",
						ReactCount: 42,
						ShareCount: 24,
						Genre: "fantasy",
					},
					{
						ID: "111111111112",
						Title: "She's Asked the Question",
						Description: "She's asked the question so many times that she barely listened to the answers anymore. The answers were always the same. Well, not exactly the same, but the same in a general sense. A more accurate description was the answers never surprised her. So, she asked for the 10,000th time, \"What's your favorite animal?\" But this time was different. When she heard the young boy's answer, she wondered if she had heard him correctly.",
						ReactCount: 256,
						ShareCount: 10,
						Genre: "lifestyle",
					},
					{
						ID: "111111111113",
						Title: "Difficult to Explain",
						Description: "It was difficult to explain to them how the diagnosis of certain death had actually given him life. While everyone around him was in tears and upset, he actually felt more at ease. The doctor said it would be less than a year. That gave him a year to live, something he'd failed to do with his daily drudgery of a routine that had passed as life until then.",
						ReactCount: 80,
						ShareCount: 100,
						Genre: "medical",
					},
					{
						ID: "111111111114",
						Title: "Sometimes",
						Description: "Sometimes that's just the way it has to be. Sure, there were probably other options, but he didn't let them enter his mind. It was done and that was that. It was just the way it had to be.",
						ReactCount: 10,
						ShareCount: 2,
						Genre: "lifestyle",
					},
					{
						ID: "111111111115",
						Title: "She Didn't Like the Food",
						Description: "She didn't like the food. She never did. She made the usual complaints and started the tantrum he knew was coming. But this time was different. Instead of trying to placate her and her unreasonable demands, he just stared at her and watched her meltdown without saying a word.",
						ReactCount: 100,
						ShareCount: 201,
						Genre: "lifestyle",
					},
					{
						ID: "111111111116",
						Title: "There was No Time",
						Description: "There was no time. He ran out of the door without half the stuff he needed for work, but it didn't matter. He was late and if he didn't make this meeting on time, someone's life may be in danger.",
						ReactCount: 5000,
						ShareCount: 250,
						Genre: "fantasy",
					},
					{
						ID: "111111111117",
						Title: "Waiting and Watching",
						Description: "Waiting and watching. It was all she had done for the past weeks. When you’re locked in a room with nothing but food and drink, that’s about all you can do anyway. She watched as birds flew past the window bolted shut. She couldn’t reach it if she wanted too, with that hole in the floor. She thought she could escape through it but three stories is a bit far down.",
						ReactCount: 323,
						ShareCount: 711,
						Genre: "motivational",
					},
					{
						ID: "111111111118",
						Title: "Their First Date",
						Description: "It was their first date and she had been looking forward to it the entire week. She had her eyes on him for months, and it had taken a convoluted scheme with several friends to make it happen, but he'd finally taken the hint and asked her out. After all the time and effort she'd invested into it, she never thought that it would be anything but wonderful. It goes without saying that things didn't work out quite as she expected.",
						ReactCount: 412,
						ShareCount: 325,
						Genre: "fantasy",
					},
					{
						ID: "111111111119",
						Title: "There was No Time",
						Description: "There was no time. He ran out of the door without half the stuff he needed for work, but it didn't matter. He was late and if he didn't make this meeting on time, someone's life may be in danger.",
						ReactCount: 5000,
						ShareCount: 250,
						Genre: "fantasy",
					},
					{
						ID: "111111111120",
						Title: "The Wolves Stopped",
						Description: "The wolves stopped in their tracks, sizing up the mother and her cubs. It had been over a week since their last meal and they were getting desperate. The cubs would make a good meal, but there were high risks taking on the mother Grizzly. A decision had to be made and the wrong choice could signal the end of the pack.",
						ReactCount: 21319,
						ShareCount: 3222,
						Genre: "fantasy",
					},
					{
						ID: "111111111121",
						Title: "The Chair Sat in the Corner",
						Description: "The chair sat in the corner where it had been for over 25 years. The only difference was there was someone actually sitting in it. How long had it been since someone had done that? Ten years or more he imagined. Yet there was no denying the presence in the chair now.",
						ReactCount: 30292,
						ShareCount: 10231,
						Genre: "horror",
					},
				},
			}
			jsonBytes, err := json.Marshal(req)
			if err != nil {
				b.FailNow()
			}

			httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, restAddr + "/big", bytes.NewBuffer(jsonBytes))
			if err != nil {
				b.FailNow()
			}
			httpReq.Header.Add("X-Country", "Indonesia")
			httpReq.Header.Add("X-Context-Id", "1bd78ba5-e59d-4ec0-8e84-709e43d14106")

			httpResp, err := s.rest.Do(httpReq)
			if err != nil {
				b.FailNow()
			}

			var resp bigResponse
			if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
				httpResp.Body.Close()
				b.FailNow()
			}

			_ = resp.Summary
			_ = resp.BookmarkedInterest
			_ = resp.OrganizedBookmark
		}
	})

	b.Run("GRPC", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req := pb.BigRequest{
				Identity: &pb.Identity{
					Username: "John Doe",
					Email: "john@doe.com",
					Country: "Indonesia",
				},
				Location: &pb.Location{
					Lat: 40.8223286,
					Lon: -96.7982002,
				},
				Interest: []string{"fantasy", "sport", "tech", "lifestyle"},
				Bookmark: []*pb.Content{
					{
						Id: "111111111111",
						Title: "Then Came the Night",
						Description: "Then came the night of the first falling star. It was seen early in the morning, rushing over Winchester eastward, a line of flame high in the atmosphere. Hundreds must have seen it and taken it for an ordinary falling star. It seemed that it fell to earth about one hundred miles east of him.",
						ReactCount: 42,
						ShareCount: 24,
						Genre: "fantasy",
					},
					{
						Id: "111111111112",
						Title: "She's Asked the Question",
						Description: "She's asked the question so many times that she barely listened to the answers anymore. The answers were always the same. Well, not exactly the same, but the same in a general sense. A more accurate description was the answers never surprised her. So, she asked for the 10,000th time, \"What's your favorite animal?\" But this time was different. When she heard the young boy's answer, she wondered if she had heard him correctly.",
						ReactCount: 256,
						ShareCount: 10,
						Genre: "lifestyle",
					},
					{
						Id: "111111111113",
						Title: "Difficult to Explain",
						Description: "It was difficult to explain to them how the diagnosis of certain death had actually given him life. While everyone around him was in tears and upset, he actually felt more at ease. The doctor said it would be less than a year. That gave him a year to live, something he'd failed to do with his daily drudgery of a routine that had passed as life until then.",
						ReactCount: 80,
						ShareCount: 100,
						Genre: "medical",
					},
					{
						Id: "111111111114",
						Title: "Sometimes",
						Description: "Sometimes that's just the way it has to be. Sure, there were probably other options, but he didn't let them enter his mind. It was done and that was that. It was just the way it had to be.",
						ReactCount: 10,
						ShareCount: 2,
						Genre: "lifestyle",
					},
					{
						Id: "111111111115",
						Title: "She Didn't Like the Food",
						Description: "She didn't like the food. She never did. She made the usual complaints and started the tantrum he knew was coming. But this time was different. Instead of trying to placate her and her unreasonable demands, he just stared at her and watched her meltdown without saying a word.",
						ReactCount: 100,
						ShareCount: 201,
						Genre: "lifestyle",
					},
					{
						Id: "111111111116",
						Title: "There was No Time",
						Description: "There was no time. He ran out of the door without half the stuff he needed for work, but it didn't matter. He was late and if he didn't make this meeting on time, someone's life may be in danger.",
						ReactCount: 5000,
						ShareCount: 250,
						Genre: "fantasy",
					},
					{
						Id: "111111111117",
						Title: "Waiting and Watching",
						Description: "Waiting and watching. It was all she had done for the past weeks. When you’re locked in a room with nothing but food and drink, that’s about all you can do anyway. She watched as birds flew past the window bolted shut. She couldn’t reach it if she wanted too, with that hole in the floor. She thought she could escape through it but three stories is a bit far down.",
						ReactCount: 323,
						ShareCount: 711,
						Genre: "motivational",
					},
					{
						Id: "111111111118",
						Title: "Their First Date",
						Description: "It was their first date and she had been looking forward to it the entire week. She had her eyes on him for months, and it had taken a convoluted scheme with several friends to make it happen, but he'd finally taken the hint and asked her out. After all the time and effort she'd invested into it, she never thought that it would be anything but wonderful. It goes without saying that things didn't work out quite as she expected.",
						ReactCount: 412,
						ShareCount: 325,
						Genre: "fantasy",
					},
					{
						Id: "111111111119",
						Title: "There was No Time",
						Description: "There was no time. He ran out of the door without half the stuff he needed for work, but it didn't matter. He was late and if he didn't make this meeting on time, someone's life may be in danger.",
						ReactCount: 5000,
						ShareCount: 250,
						Genre: "fantasy",
					},
					{
						Id: "111111111120",
						Title: "The Wolves Stopped",
						Description: "The wolves stopped in their tracks, sizing up the mother and her cubs. It had been over a week since their last meal and they were getting desperate. The cubs would make a good meal, but there were high risks taking on the mother Grizzly. A decision had to be made and the wrong choice could signal the end of the pack.",
						ReactCount: 21319,
						ShareCount: 3222,
						Genre: "fantasy",
					},
					{
						Id: "111111111121",
						Title: "The Chair Sat in the Corner",
						Description: "The chair sat in the corner where it had been for over 25 years. The only difference was there was someone actually sitting in it. How long had it been since someone had done that? Ten years or more he imagined. Yet there was no denying the presence in the chair now.",
						ReactCount: 30292,
						ShareCount: 10231,
						Genre: "horror",
					},
				},
			}

			headers := metadata.New(map[string]string{
				"X-Country": "Indonesia",
				"X-Context-Id": "1bd78ba5-e59d-4ec0-8e84-709e43d14106",
			})
			reqCtx := metadata.NewOutgoingContext(ctx, headers)
			resp, err := s.grpc.Big(reqCtx, &req)
			if err != nil {
				b.FailNow()
			}

			_ = resp.GetSummary()
			_ = resp.GetBookmarkedInterest()
			_ = resp.GetOrganizedBookmark()
		}
	})
}