package client

import (
	"fmt"
	"path"
	"strings"

	"github.com/ligato/cn-infra/datasync"
	"github.com/ligato/cn-infra/db/keyval"
	"github.com/ligato/cn-infra/logging"
	"github.com/ligato/cn-infra/servicelabel"
)

// KVDBClient provides client access to the KVDB server.
type KVDBClient struct {
	keyval.CoreBrokerWatcher
	serviceLabel string
}

func NewKVDBClient(kvdb keyval.CoreBrokerWatcher, serviceLabel string) *KVDBClient {
	return &KVDBClient{
		CoreBrokerWatcher: kvdb,
		serviceLabel:      serviceLabel,
	}
}

func (k *KVDBClient) Put(key string, data []byte, opts ...datasync.PutOption) (err error) {
	key, err = k.completeFullKey(key)
	if err != nil {
		return err
	}
	logging.Debugf("KVDBClient.Put: %s", key)

	return k.CoreBrokerWatcher.Put(key, data, opts...)
}

func (k *KVDBClient) GetValue(key string) (data []byte, found bool, revision int64, err error) {
	key, err = k.completeFullKey(key)
	if err != nil {
		return nil, false, 0, err
	}
	logging.Debugf("KVDBClient.GetValue: %s", key)

	return k.CoreBrokerWatcher.GetValue(key)
}

func (k *KVDBClient) ListValues(prefix string) (keyval.BytesKeyValIterator, error) {
	prefix = ensureAllAgentsPrefix(prefix)
	logging.Debugf("KVDBClient.ListValues: %s", prefix)

	return k.CoreBrokerWatcher.ListValues(prefix)
}

func (k *KVDBClient) ListKeys(prefix string) (keyval.BytesKeyIterator, error) {
	prefix = ensureAllAgentsPrefix(prefix)
	logging.Debugf("KVDBClient.ListKeys: %s", prefix)

	return k.CoreBrokerWatcher.ListKeys(prefix)
}

func (k *KVDBClient) Delete(key string, opts ...datasync.DelOption) (existed bool, err error) {
	key, err = k.completeFullKey(key)
	if err != nil {
		return false, err
	}
	logging.Debugf("KVDBClient.Delete: %s", key)

	return k.CoreBrokerWatcher.Delete(key, opts...)
}

func (k *KVDBClient) completeFullKey(key string) (string, error) {
	if strings.HasPrefix(key, servicelabel.GetAllAgentsPrefix()) {
		return key, nil
	}
	if k.serviceLabel == "" {
		return "", fmt.Errorf("service label is not defined, cannot get complete key")
	}
	key = path.Join(servicelabel.GetAllAgentsPrefix(), k.serviceLabel, key)
	return key, nil
}

func ensureAllAgentsPrefix(key string) string {
	if strings.HasPrefix(key, servicelabel.GetAllAgentsPrefix()) {
		return key
	}
	return path.Join(servicelabel.GetAllAgentsPrefix(), key)
}
