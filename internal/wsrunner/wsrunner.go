// Package wsrunner は WebSocket サンプルで共通して使う購読ループを提供する。
//
// gogmocoin の WebSocket クライアントはいずれも Subscribe / Stream / Unsubscribe を
// 備えており、サンプルごとに「購読 → 受信ループ → 購読解除」のコードがほぼ同一になる。
// その重複と、Ctrl-C 時に Unsubscribe が呼ばれない問題をまとめて解消するためのヘルパー。
package wsrunner

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"
)

const (
	// idleTimeout は 1 回の受信待ちで許容する無通信時間。
	idleTimeout = 180 * time.Second
	// maxIdleCount は idleTimeout が連続したときに終了するまでの回数。
	maxIdleCount = 20
)

// Run はストリームを購読し、受信メッセージをログ出力し続ける。
// 以下のいずれかが起きると Unsubscribe を呼んで終了する:
//   - SIGINT / SIGTERM（Ctrl-C 等）を受信したとき
//   - idleTimeout が maxIdleCount 回連続したとき
//
// stream は「受信チャネルを返す関数」（クライアントの Stream() ）を渡す。
// Stream() は Subscribe 中はチャネルを内部でメモ化して同じものを返すが、
// このヘルパーでも Subscribe 後に一度だけ呼んで使い回す。
func Run[T any](subscribe func() error, stream func() <-chan T, unsubscribe func() error) error {
	if err := subscribe(); err != nil {
		return err
	}
	// 正常終了・シグナル終了のどちらでも購読解除して接続をクリーンに閉じる。
	defer func() {
		if err := unsubscribe(); err != nil {
			log.Printf("unsubscribe failed: %v", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ch := stream()
	idleCount := 0
	for {
		select {
		case <-ctx.Done():
			log.Println("signal received, shutting down...")
			return nil
		case v := <-ch:
			log.Printf("msg:%+v\n", v)
		case <-time.After(idleTimeout):
			log.Println("timeout...")
			idleCount++
			if idleCount >= maxIdleCount {
				log.Println("idle limit reached, shutting down...")
				return nil
			}
		}
	}
}
