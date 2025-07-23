package storage

type InMemoryStorage struct {
	mu    sync.RWMutex
	store map[string]string
}


func NewInMemoryStorage(expiration time.Duration) *InMemoryStorage{
	return &InMemoryStorage{store: make(map[string)string}
}

func (mem *InMemoryStorage) Set(ctx context.Context, key, value string) error {

}

func (mem *InMemoryStorage) Get(ctx context.Context, key string) (string, error){

}