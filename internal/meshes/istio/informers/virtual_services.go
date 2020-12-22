package informers

import (
	"log"

	"github.com/layer5io/meshsync/pkg/model"
	broker "github.com/layer5io/meshsync/pkg/broker"
	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	"k8s.io/client-go/tools/cache"
)

func (i *Istio) VirtualServiceInformer() cache.SharedIndexInformer {
	// get informer
	VirtualServiceInformer := i.client.GetVirtualServiceInformer().Informer()

	// register event handlers
	VirtualServiceInformer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				VirtualService := obj.(*v1beta1.VirtualService)
				log.Printf("VirtualService Named: %s - added", VirtualService.Name)
				err := i.broker.Publish(Subject, &broker.Message{
					Object: model.ConvObject(
						VirtualService.TypeMeta,
						VirtualService.ObjectMeta,
						VirtualService.Spec,
						VirtualService.Status,
					)})
				if err != nil {
					log.Println("NATS: Error publishing VirtualService")
				}
			},
			UpdateFunc: func(new interface{}, old interface{}) {
				VirtualService := new.(*v1beta1.VirtualService)
				log.Printf("VirtualService Named: %s - updated", VirtualService.Name)
				err := i.broker.Publish(Subject, &broker.Message{
					Object: model.ConvObject(
						VirtualService.TypeMeta,
						VirtualService.ObjectMeta,
						VirtualService.Spec,
						VirtualService.Status,
					)})
				if err != nil {
					log.Println("NATS: Error publishing VirtualService")
				}
			},
			DeleteFunc: func(obj interface{}) {
				VirtualService := obj.(*v1beta1.VirtualService)
				log.Printf("VirtualService Named: %s - deleted", VirtualService.Name)
				err := i.broker.Publish(Subject, &broker.Message{
					Object: model.ConvObject(
						VirtualService.TypeMeta,
						VirtualService.ObjectMeta,
						VirtualService.Spec,
						VirtualService.Status,
					)})
				if err != nil {
					log.Println("NATS: Error publishing VirtualService")
				}
			},
		},
	)

	return VirtualServiceInformer
}
