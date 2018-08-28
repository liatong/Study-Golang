// demo6mlib project demo6mlib.go
package demo6mlib

import (
	"errors"
	"fmt"
	"time"
)

//编程思路整理。
//需求： 要做一个音乐管理器，那就需要先构造一个结构，用来当作一首歌曲实例。
//另外构造一个音乐播放器，能够对歌曲做增删改查。

//step1 先构造一个音乐介质出来。

//然后还有不同对音乐体制，不同对音乐体制播放对方式不一样。

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	music map[string]MusicEntry
	id    string
}

func NewMusicManager() *MusicManager {
	//为什么要在这里返回的是一个指针。因为我在new方法里面构造的，就需要把这个对象原本的传递回去。
	return &MusicManager{make(map[string]MusicEntry), "id001"}
}

func (m *MusicManager) Len() int {
	return len(m.music)
}

func (m *MusicManager) Get(name string) (music *MusicEntry, err error) {
	//	这里需要注意map的获取方式要对是通过 map[key] 来进行获取的。
	//	if music, ok := m.music（）; ok {
	if music, ok := m.music[name]; ok {
		return &music, nil
	}
	return nil, errors.New("nomusic")
}
func (m *MusicManager) Add(mu *MusicEntry) {
	//既然传递进来的是一个指针，你返回给人家的也是一个指针，那么就要存到map里面也应该是一个指针才是。
	m.music[mu.Name] = *mu
}
func (m *MusicManager) Del(name string) {
	delete(m.music, name)
}

//除此之外还需要构造一个播放器，能够实现，歌曲对录入，歌曲对播放，歌曲对暂停等动作等执行。主要负责和人机交互等部分。

type Player struct {
	//这里我需要的是一个
	musicm MusicManager
}

func (p *Player) Play(name string) {
	music, err := p.musicm.Get(name)
	if err != nil {
		fmt.Print("not have music ")
		return
	}

	switch music.Type {
	case "MP3":
		p := &MP3Player{}
		p.Play(music.Source)
	case "MP4":
		p := &MP4Player{}
		p.Play(music.Source)
	default:
		fmt.Print("default player nothing to do...")
		return
	}

}

type MP3Player struct {
	state string
}

func (mp3 *MP3Player) Play(source string) {

	fmt.Print("mp3 playing..", source)
	time.Sleep(1000 * time.Millisecond)

}

type MP4Player struct {
	state string
}

func (mp4 MP4Player) Play(source string) {
	fmt.Print("mp4 playing ...", source)
	time.Sleep(10000 * time.Millisecond)
}
